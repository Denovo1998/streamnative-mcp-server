// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/streamnative/streamnative-mcp-server/pkg/common"
	mcppulsar "github.com/streamnative/streamnative-mcp-server/pkg/pulsar"
)

// PulsarClientAddConsumerTools adds Pulsar client consumer tools to the MCP server
func PulsarClientAddConsumerTools(s *server.MCPServer, _ bool, features []string) {
	if !slices.Contains(features, string(FeaturePulsarClient)) && !slices.Contains(features, string(FeatureAll)) && !slices.Contains(features, string(FeatureAllPulsar)) {
		return
	}

	// Main consume tool
	consumeTool := mcp.NewTool("pulsar_client_consume",
		mcp.WithDescription("Consume messages from a Pulsar topic. This tool allows you to consume messages "+
			"from a specified Pulsar topic with various options to control the subscription behavior, "+
			"message processing, and display format. Do not use this tool for Kafka protocol operations. Use 'kafka_client_consume' instead."),
		mcp.WithString("topic", mcp.Required(),
			mcp.Description("Topic to consume from"),
		),
		mcp.WithString("subscription-name", mcp.Required(),
			mcp.Description("Subscription name"),
		),
		mcp.WithString("subscription-type",
			mcp.Description("Subscription type: exclusive, shared, failover, key_shared (default: exclusive)"),
		),
		mcp.WithString("subscription-mode",
			mcp.Description("Subscription mode: durable, non-durable (default: durable)"),
		),
		mcp.WithString("initial-position",
			mcp.Description("Initial position: latest (consume from the latest message)"+
				", earliest (consume from the earliest message) (default: latest)"),
		),
		mcp.WithNumber("num-messages",
			mcp.Description("Number of messages to consume (default: 10)"),
		),
		mcp.WithNumber("timeout",
			mcp.Description("Timeout for consuming messages in seconds (default: 30)"),
		),
		mcp.WithBoolean("show-properties",
			mcp.Description("Show message properties (default: false)"),
		),
		mcp.WithBoolean("hide-payload",
			mcp.Description("Hide message payload (default: false)"),
		),
	)
	s.AddTool(consumeTool, handleClientConsume)
}

func handleClientConsume(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract required parameters with validation
	topic, err := common.RequiredParam[string](request.Params.Arguments, "topic")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get topic: %v", err)), nil
	}

	subscriptionName, err := common.RequiredParam[string](request.Params.Arguments, "subscription-name")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get subscription name: %v", err)), nil
	}

	// Set default values and extract optional parameters
	subscriptionType := "exclusive"
	if val, exists := common.OptionalParam[string](request.Params.Arguments, "subscription-type"); exists && val != "" {
		subscriptionType = val
	}

	subscriptionMode := "durable"
	if val, exists := common.OptionalParam[string](request.Params.Arguments, "subscription-mode"); exists && val != "" {
		subscriptionMode = val
	}

	initialPosition := "latest"
	if val, exists := common.OptionalParam[string](request.Params.Arguments, "initial-position"); exists && val != "" {
		initialPosition = val
	}

	numMessages := 10
	if val, exists := common.OptionalParam[float64](request.Params.Arguments, "num-messages"); exists {
		numMessages = int(val)
	}

	timeout := 30
	if val, exists := common.OptionalParam[float64](request.Params.Arguments, "timeout"); exists {
		timeout = int(val)
	}

	showProperties := false
	if val, exists := common.OptionalParam[bool](request.Params.Arguments, "show-properties"); exists {
		showProperties = val
	}

	hidePayload := false
	if val, exists := common.OptionalParam[bool](request.Params.Arguments, "hide-payload"); exists {
		hidePayload = val
	}

	// Setup client
	client, err := mcppulsar.GetPulsarClient()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to create Pulsar client: %v", err)), nil
	}
	defer client.Close()

	// Prepare consumer options
	consumerOpts := pulsar.ConsumerOptions{
		Name:             "snmcp-consumer",
		SubscriptionName: subscriptionName,
	}

	consumerOpts.Topic = topic

	// Set subscription type
	switch strings.ToLower(subscriptionType) {
	case "exclusive":
		consumerOpts.Type = pulsar.Exclusive
	case "shared":
		consumerOpts.Type = pulsar.Shared
	case "failover":
		consumerOpts.Type = pulsar.Failover
	case "key_shared":
		consumerOpts.Type = pulsar.KeyShared
	default:
		return mcp.NewToolResultError(fmt.Sprintf("Invalid subscription type: %s", subscriptionType)), nil
	}

	// Set subscription mode
	switch strings.ToLower(subscriptionMode) {
	case "durable":
		consumerOpts.SubscriptionMode = pulsar.Durable
	case "non-durable":
		consumerOpts.SubscriptionMode = pulsar.NonDurable
	default:
		return mcp.NewToolResultError(fmt.Sprintf("Invalid subscription mode: %s", subscriptionMode)), nil
	}

	// Set initial position
	switch strings.ToLower(initialPosition) {
	case "latest":
		consumerOpts.SubscriptionInitialPosition = pulsar.SubscriptionPositionLatest
	case "earliest":
		consumerOpts.SubscriptionInitialPosition = pulsar.SubscriptionPositionEarliest
	default:
		return mcp.NewToolResultError(fmt.Sprintf("Invalid initial position: %s", initialPosition)), nil
	}

	// Create consumer
	consumer, err := client.Subscribe(consumerOpts)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to create consumer: %v", err)), nil
	}
	defer consumer.Close()

	// Set up timeout context
	timeoutDuration := time.Duration(timeout) * time.Second
	consumeCtx, cancelConsume := context.WithTimeout(ctx, timeoutDuration)
	defer cancelConsume()

	// Container for messages
	type MessageData struct {
		ID           string            `json:"id"`
		PublishTime  string            `json:"publish_time"`
		Properties   map[string]string `json:"properties,omitempty"`
		Key          string            `json:"key,omitempty"`
		Data         string            `json:"data,omitempty"`
		MessageCount int               `json:"message_count"`
	}

	messages := []MessageData{}
	messageCount := 0

	// Consume messages
	for {
		// Check if we've consumed the requested number of messages
		if numMessages > 0 && messageCount >= numMessages {
			break
		}

		// Receive message with timeout
		msg, err := consumer.Receive(consumeCtx)
		if err != nil {
			if err == context.DeadlineExceeded {
				break
			}
			if err == context.Canceled {
				break
			}
			return mcp.NewToolResultError(fmt.Sprintf("Error receiving message: %v", err)), nil
		}

		// Process the message
		messageCount++

		// Create message data
		messageData := MessageData{
			ID:           msg.ID().String(),
			PublishTime:  msg.PublishTime().Format(time.RFC3339),
			MessageCount: messageCount,
		}

		// Add properties if requested
		if showProperties {
			messageData.Properties = msg.Properties()
		}

		// Add key if present
		if msg.Key() != "" {
			messageData.Key = msg.Key()
		}

		// Add payload unless hidden
		if !hidePayload {
			messageData.Data = string(msg.Payload())
		}

		messages = append(messages, messageData)

		// Acknowledge the message
		_ = consumer.Ack(msg)
	}

	// Prepare response
	response := map[string]interface{}{
		"topic":             topic,
		"subscription_name": subscriptionName,
		"messages_consumed": messageCount,
		"messages":          messages,
	}

	// Convert to JSON
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to encode result: %v", err)), nil
	}

	return mcp.NewToolResultText(string(jsonBytes)), nil
}

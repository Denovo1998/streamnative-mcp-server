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

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/streamnative/streamnative-mcp-server/pkg/common"
	"github.com/streamnative/streamnative-mcp-server/pkg/kafka"
	"github.com/twmb/franz-go/pkg/kadm"
)

func KafkaAdminAddGroupsTools(s *server.MCPServer, readOnly bool, features []string) {
	if !slices.Contains(features, string(FeatureKafkaAdmin)) && !slices.Contains(features, string(FeatureAll)) && !slices.Contains(features, string(FeatureAllKafka)) {
		return
	}

	resourceDesc := "Resource to operate on. Available resources:\n" +
		"- group: A single Kafka Consumer Group for operations on individual groups (describe, remove-members, set-offset, delete-offset)\n" +
		"- groups: Collection of Kafka Consumer Groups for bulk operations (list)"

	operationDesc := "Operation to perform. Available operations:\n" +
		"- list: List all Kafka Consumer Groups in the cluster\n" +
		"- describe: Get detailed information about a specific Consumer Group, including members, offsets, and lag\n" +
		"- remove-members: Remove specific members from a Consumer Group to force rebalancing or troubleshoot issues\n" +
		"- offsets: Get offsets for a specific consumer group\n" +
		"- delete-offset: Delete a specific offset for a consumer group of a topic\n" +
		"- set-offset: Set a specific offset for a consumer group's topic-partition"

	toolDesc := "Unified tool for managing Apache Kafka Consumer Groups.\n" +
		"This tool provides access to Kafka consumer group operations including listing, describing, and managing group membership.\n" +
		"Kafka Consumer Groups are a key concept for scalable consumption of Kafka topics. A consumer group consists of multiple consumer instances\n" +
		"that collaborate to consume data from topic partitions. Kafka ensures that:\n" +
		"- Each partition is consumed by exactly one consumer in the group\n" +
		"- When consumers join or leave, Kafka triggers a 'rebalance' to redistribute partitions\n" +
		"- Consumer groups track consumption progress through committed offsets\n\n" +
		"Usage Examples:\n\n" +
		"1. List all Kafka Consumer Groups in the cluster:\n" +
		"   resource: \"groups\"\n" +
		"   operation: \"list\"\n\n" +
		"2. Describe a specific Kafka Consumer Group to see its members and consumption details:\n" +
		"   resource: \"group\"\n" +
		"   operation: \"describe\"\n" +
		"   group: \"my-consumer-group\"\n\n" +
		"3. Remove specific members from a Kafka Consumer Group to trigger rebalancing:\n" +
		"   resource: \"group\"\n" +
		"   operation: \"remove-members\"\n" +
		"   group: \"my-consumer-group\"\n" +
		"   members: \"consumer-instance-1,consumer-instance-2\"\n\n" +
		"4. Get offsets for a specific consumer group:\n" +
		"   resource: \"group\"\n" +
		"   operation: \"offsets\"\n" +
		"   group: \"my-consumer-group\"\n\n" +
		"5. Delete a specific offset for a consumer group of a topic:\n" +
		"   resource: \"group\"\n" +
		"   operation: \"delete-offset\"\n" +
		"   group: \"my-consumer-group\"\n" +
		"   topic: \"my-topic\"\n\n" +
		"6. Set a specific offset for a consumer group's topic-partition:\n" +
		"   resource: \"group\"\n" +
		"   operation: \"set-offset\"\n" +
		"   group: \"my-consumer-group\"\n" +
		"   topic: \"my-topic\"\n" +
		"   partition: 0\n" +
		"   offset: 1000\n\n" +
		"This tool requires Kafka super-user permissions."

	kafkaGroupsTool := mcp.NewTool("kafka_admin_groups",
		mcp.WithDescription(toolDesc),
		mcp.WithString("resource", mcp.Required(),
			mcp.Description(resourceDesc),
		),
		mcp.WithString("operation", mcp.Required(),
			mcp.Description(operationDesc),
		),
		mcp.WithString("group",
			mcp.Description("The name of the Kafka Consumer Group to operate on. "+
				"Required for the 'describe' and 'remove-members' operations. "+
				"Must be an existing consumer group name in the Kafka cluster. "+
				"Consumer Group names are case-sensitive and typically follow a naming convention like 'application-name'.")),
		mcp.WithString("members",
			mcp.Description("The members to remove from the Kafka Consumer Group. "+
				"Required for the 'remove-members' operation. "+
				"Specify as a comma-separated list of member instance IDs (no spaces). "+
				"Removing members will trigger a consumer group rebalance. "+
				"Example: 'consumer-instance-1,consumer-instance-2'.")),
		mcp.WithString("topic",
			mcp.Description("The name of the Kafka topic to operate on. "+
				"Required for the 'delete' operation. "+
				"Must be an existing topic name in the Kafka cluster.")),
		mcp.WithNumber("partition",
			mcp.Description("The partition number to set offset for. "+
				"Required for the 'set-offset' operation. "+
				"Must be a valid partition number for the specified topic.")),
		mcp.WithNumber("offset",
			mcp.Description("The offset value to set. "+
				"Required for the 'set-offset' operation. "+
				"Use -1 for earliest offset, -2 for latest offset, or a specific positive value.")),
	)

	s.AddTool(kafkaGroupsTool, handleKafkaGroupsTool(readOnly))
}

func handleKafkaGroupsTool(readOnly bool) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Get required parameters
		resource, err := common.RequiredParam[string](request.Params.Arguments, "resource")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get resource: %v", err)), nil
		}

		operation, err := common.RequiredParam[string](request.Params.Arguments, "operation")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get operation: %v", err)), nil
		}

		// Normalize parameters
		resource = strings.ToLower(resource)
		operation = strings.ToLower(operation)

		// Validate write operations in read-only mode
		if readOnly && (operation == "remove-members" || operation == "delete-offset") {
			return mcp.NewToolResultError("Write operations are not allowed in read-only mode"), nil
		}

		// Create the admin client
		admin, err := kafka.GetKafkaAdminClient()
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get admin client: %v", err)), nil
		}

		// Dispatch based on resource and operation
		switch resource {
		case "group":
			switch operation {
			case "describe":
				return handleKafkaGroupDescribe(ctx, admin, request)
			case "remove-members":
				return handleKafkaGroupRemoveMembers(ctx, admin, request)
			case "offsets":
				return handleKafkaGroupOffsets(ctx, admin, request)
			case "delete-offset":
				return handleKafkaGroupDeleteOffset(ctx, admin, request)
			case "set-offset":
				return handleKafkaGroupSetOffset(ctx, admin, request)
			default:
				return mcp.NewToolResultError(fmt.Sprintf("Invalid operation for resource 'group': %s", operation)), nil
			}
		case "groups":
			switch operation {
			case "list":
				return handleKafkaGroupsList(ctx, admin, request)
			default:
				return mcp.NewToolResultError(fmt.Sprintf("Invalid operation for resource 'groups': %s", operation)), nil
			}
		default:
			return mcp.NewToolResultError(fmt.Sprintf("Invalid resource: %s. Available resources: group, groups", resource)), nil
		}
	}
}

func handleKafkaGroupDescribe(ctx context.Context, admin *kadm.Client, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	groupName, err := common.RequiredParam[string](request.Params.Arguments, "group")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get group name: %v", err)), nil
	}

	response, err := admin.DescribeGroups(ctx, groupName)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to describe group: %v", err)), nil
	}

	lags, err := admin.Lag(ctx, groupName)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get lag: %v", err)), nil
	}

	result := map[string]interface{}{
		"group": response,
		"lag":   lags,
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal response: %v", err)), nil
	}

	return mcp.NewToolResultText(string(jsonBytes)), nil
}

func handleKafkaGroupRemoveMembers(ctx context.Context, admin *kadm.Client, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	groupName, err := common.RequiredParam[string](request.Params.Arguments, "group")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get group name: %v", err)), nil
	}

	members, err := common.RequiredParam[string](request.Params.Arguments, "members")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get members: %v", err)), nil
	}
	builder := kadm.LeaveGroup(groupName).InstanceIDs(strings.Split(members, ",")...)
	response, err := admin.LeaveGroup(ctx, builder)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to remove members: %v", err)), nil
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal response: %v", err)), nil
	}

	return mcp.NewToolResultText(string(jsonBytes)), nil
}

func handleKafkaGroupsList(ctx context.Context, admin *kadm.Client, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	response, err := admin.ListGroups(ctx)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to list groups: %v", err)), nil
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal response: %v", err)), nil
	}

	return mcp.NewToolResultText(string(jsonBytes)), nil
}

func handleKafkaGroupOffsets(ctx context.Context, admin *kadm.Client, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	groupName, err := common.RequiredParam[string](request.Params.Arguments, "group")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get group name: %v", err)), nil
	}

	response, err := admin.FetchOffsets(ctx, groupName)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to list offsets: %v", err)), nil
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal response: %v", err)), nil
	}

	return mcp.NewToolResultText(string(jsonBytes)), nil
}

func handleKafkaGroupDeleteOffset(ctx context.Context, admin *kadm.Client, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	groupName, err := common.RequiredParam[string](request.Params.Arguments, "group")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get group name: %v", err)), nil
	}

	topic, err := common.RequiredParam[string](request.Params.Arguments, "topic")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get topic name: %v", err)), nil
	}

	// Create a TopicsSet with the specified topic
	// This will target all partitions for the given topic
	topicsSet := make(kadm.TopicsSet)
	topicsSet.Add(topic)

	// Call DeleteOffsets to delete the offsets for the specified topic
	responses, err := admin.DeleteOffsets(ctx, groupName, topicsSet)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to delete offsets: %v", err)), nil
	}

	// Check for errors in the response
	if err := responses.Error(); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Errors occurred during offset deletion: %v", err)), nil
	}

	jsonBytes, err := json.Marshal(responses)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal response: %v", err)), nil
	}

	return mcp.NewToolResultText(string(jsonBytes)), nil
}

func handleKafkaGroupSetOffset(ctx context.Context, admin *kadm.Client, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Get required parameters
	groupName, err := common.RequiredParam[string](request.Params.Arguments, "group")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get group name: %v", err)), nil
	}

	topic, err := common.RequiredParam[string](request.Params.Arguments, "topic")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get topic name: %v", err)), nil
	}

	partition, err := common.RequiredParam[float64](request.Params.Arguments, "partition")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get partition number: %v", err)), nil
	}
	partitionInt := int32(partition)

	offset, err := common.RequiredParam[float64](request.Params.Arguments, "offset")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get offset value: %v", err)), nil
	}
	offsetInt := int64(offset)

	// Create Offsets object with the specified topic, partition, and offset
	offsets := make(kadm.Offsets)
	offsets.AddOffset(topic, partitionInt, offsetInt, -1) // Using -1 for leaderEpoch as it's optional

	// Commit the offsets
	responses, err := admin.CommitOffsets(ctx, groupName, offsets)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to commit offsets: %v", err)), nil
	}

	// Check for errors in the response
	if err := responses.Error(); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Errors occurred during offset commit: %v", err)), nil
	}

	jsonBytes, err := json.Marshal(responses)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal response: %v", err)), nil
	}

	return mcp.NewToolResultText(string(jsonBytes)), nil
}

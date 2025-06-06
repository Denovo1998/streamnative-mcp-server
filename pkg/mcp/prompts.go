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

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/streamnative/streamnative-mcp-server/pkg/common"
	"github.com/streamnative/streamnative-mcp-server/pkg/config"
	sncloud "github.com/streamnative/streamnative-mcp-server/sdk/sdk-apiserver"
	"k8s.io/utils/ptr"
)

type ServerlessPoolMember struct {
	Provider  string
	Namespace string
	Pool      string
	Location  string
}

var (
	ServerlessPoolMemberList = []ServerlessPoolMember{
		{
			Provider:  "azure",
			Namespace: "streamnative",
			Pool:      "shared-azure",
			Location:  "eastus",
		},
		{
			Provider:  "aws",
			Namespace: "streamnative",
			Pool:      "shared-aws",
			Location:  "us-east-2",
		},
		// {
		// 	Provider:  "gcloud",
		// 	Namespace: "streamnative",
		// 	Pool:      "shared-gcp",
		// 	Location:  "us-central1",
		// },
	}
	AvailableProviders = []string{"azure", "aws", "gcloud"}
)

func RegisterPrompts(s *server.MCPServer) {
	s.AddPrompt(mcp.NewPrompt("list-sncloud-clusters",
		mcp.WithPromptDescription("List all clusters from the StreamNative Cloud"),
	), handleListPulsarClusters)
	s.AddPrompt(mcp.NewPrompt("read-sncloud-cluster",
		mcp.WithPromptDescription("Read a cluster from the StreamNative Cloud"),
		mcp.WithArgument("name", mcp.RequiredArgument(), mcp.ArgumentDescription("The name of the cluster")),
	), handleReadPulsarCluster)
	s.AddPrompt(
		mcp.NewPrompt("build-sncloud-serverless-cluster",
			mcp.WithPromptDescription("Build a Serverless cluster in the StreamNative Cloud"),
			mcp.WithArgument("instance-name", mcp.RequiredArgument(), mcp.ArgumentDescription("The name of the Pulsar instance, cannot reuse the name of existing instance.")),
			mcp.WithArgument("cluster-name", mcp.RequiredArgument(), mcp.ArgumentDescription("The name of the Pulsar cluster, cannot reuse the name of existing cluster.")),
			mcp.WithArgument("provider", mcp.ArgumentDescription("The cloud provider, could be `aws`, `gcp`, `azure`. If the selected provider do not serve serverless cluster, the prompt will return an error. If not specified, the system will use a random provider depending on the availability.")),
		),
		handleBuildServerlessPulsarCluster,
	)
}

func handleListPulsarClusters(ctx context.Context, _ mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	options := common.GetOptions(ctx)
	apiClient, err := config.GetAPIClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get API client: %v", err)
	}

	clusters, clustersBody, err := apiClient.CloudStreamnativeIoV1alpha1Api.ListCloudStreamnativeIoV1alpha1NamespacedPulsarCluster(ctx, options.Organization).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to list pulsar clusters: %v", err)
	}
	defer clustersBody.Body.Close()

	var messages = make(
		[]mcp.PromptMessage,
		len(clusters.Items)+1,
	)

	messages[0] = mcp.PromptMessage{
		Content: mcp.TextContent{
			Type: "text",
			Text: fmt.Sprintf(
				"There are %d Pulsar clusters in the StreamNative Cloud from organization %s:",
				len(clusters.Items),
				options.Organization,
			),
		},
		Role: mcp.RoleUser,
	}

	for i, cluster := range clusters.Items {
		instanceName := cluster.Spec.InstanceName
		displayName := cluster.Spec.DisplayName
		if displayName == nil || *displayName == "" {
			displayName = cluster.Metadata.Name
		}

		status := "Not Ready"
		if common.IsClusterAvailable(cluster) {
			status = "Ready"
		}

		engineType := common.GetEngineType(cluster)

		messages[i+1] = mcp.PromptMessage{
			Content: mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf(
					"Instance Name: %s\nCluster Name: %s\nCluster Display Name: %s\nCluster Status: %s\nCluster Engine Type: %s",
					instanceName,
					*cluster.Metadata.Name,
					*displayName,
					status,
					engineType,
				),
			},
			Role: mcp.RoleUser,
		}
	}

	return &mcp.GetPromptResult{
		Description: fmt.Sprintf("Pulsar clusters from StreamNative Cloud organization %s, you can use `sncloud_context_use_cluster` tool to switch to selected cluster, and use pulsar and kafka tools to interact with the cluster.", options.Organization),
		Messages:    messages,
	}, nil
}

func handleReadPulsarCluster(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	options := common.GetOptions(ctx)
	apiClient, err := config.GetAPIClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get API client: %v", err)
	}

	name, err := common.RequiredParam[string](common.ConvertToMapInterface(request.Params.Arguments), "name")
	if err != nil {
		return nil, fmt.Errorf("failed to get name: %v", err)
	}

	clusters, clustersBody, err := apiClient.CloudStreamnativeIoV1alpha1Api.ListCloudStreamnativeIoV1alpha1NamespacedPulsarCluster(ctx, options.Organization).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to list pulsar clusters: %v", err)
	}
	defer clustersBody.Body.Close()
	var cluster sncloud.ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarCluster
	for _, c := range clusters.Items {
		if *c.Metadata.Name == name {
			cluster = c
			break
		}
	}

	if cluster.Metadata == nil {
		return nil, fmt.Errorf("failed to find pulsar cluster: %s", name)
	}

	if cluster.Metadata != nil && len(cluster.Metadata.ManagedFields) > 0 {
		cluster.Metadata.ManagedFields = nil
	}

	context, err := json.Marshal(cluster)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal cluster: %v", err)
	}

	var messages = make(
		[]mcp.PromptMessage,
		1,
	)

	messages[0] = mcp.PromptMessage{
		Content: mcp.TextContent{
			Type: "text",
			Text: string(context),
		},
		Role: mcp.RoleUser,
	}

	return &mcp.GetPromptResult{
		Description: fmt.Sprintf("Detailed information of Pulsar cluster %s, you can use `sncloud_context_use_cluster` tool to switch to this cluster, and use pulsar and kafka tools to interact with the cluster.", name),
		Messages:    messages,
	}, nil
}

func handleBuildServerlessPulsarCluster(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	options := common.GetOptions(ctx)
	apiClient, err := config.GetAPIClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get API client: %v", err)
	}
	arguments := common.ConvertToMapInterface(request.Params.Arguments)

	instanceName, err := common.RequiredParam[string](arguments, "instance-name")
	if err != nil {
		return nil, fmt.Errorf("failed to get instance name: %v", err)
	}

	clusterName, err := common.RequiredParam[string](arguments, "cluster-name")
	if err != nil {
		return nil, fmt.Errorf("failed to get cluster name: %v", err)
	}

	provider, hasProvider := common.OptionalParam[string](arguments, "provider")
	if !hasProvider {
		provider = ""
	}
	if provider != "" {
		if !slices.Contains(AvailableProviders, provider) {
			return nil, fmt.Errorf("invalid provider: %s, available providers: %v", provider, AvailableProviders)
		}
	}

	poolOptions, poolOptionsBody, err := apiClient.CloudStreamnativeIoV1alpha1Api.ListCloudStreamnativeIoV1alpha1NamespacedPoolOption(ctx, options.Organization).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to list pool options: %v", err)
	}
	defer poolOptionsBody.Body.Close()
	if poolOptions == nil {
		return nil, fmt.Errorf("no pool options found")
	}

	var poolRef *sncloud.ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolRef
	var selectedLocation *string

	for _, poolOpt := range poolOptions.Items {
		if pr, ok := poolOpt.Spec.GetPoolRefOk(); ok {
			for _, poolMember := range ServerlessPoolMemberList {
				if provider != "" && poolOpt.Spec.CloudType != provider {
					continue
				}
				if pr.Name == poolMember.Pool && pr.Namespace == poolMember.Namespace {
					for _, location := range poolOpt.Spec.Locations {
						if location.Location == poolMember.Location {
							poolRef = pr
							selectedLocation = &location.Location
							break
						}
					}
				}
			}
		}
	}

	if poolRef == nil || selectedLocation == nil {
		return nil, fmt.Errorf("no available pool")
	}

	inst := sncloud.ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarInstance{}
	clus := sncloud.ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarCluster{}

	inst.ApiVersion = ptr.To("cloud.streamnative.io/v1alpha1")
	inst.Kind = ptr.To("PulsarInstance")
	inst.Metadata = &sncloud.V1ObjectMeta{
		Name:      &instanceName,
		Namespace: &options.Organization,
		Labels: &map[string]string{
			"managed-by": "streamnative-mcp",
		},
	}

	inst.Spec = &sncloud.ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarInstanceSpec{
		AvailabilityMode: "zonal",
		PoolRef:          poolRef,
		Type:             ptr.To("serverless"),
	}

	clus.ApiVersion = ptr.To("cloud.streamnative.io/v1alpha1")
	clus.Kind = ptr.To("PulsarCluster")
	clus.Metadata = &sncloud.V1ObjectMeta{
		Name:      ptr.To(""),
		Namespace: &options.Organization,
		Labels: &map[string]string{
			"managed-by": "streamnative-mcp",
		},
	}

	clus.Spec = &sncloud.ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarClusterSpec{
		Broker: sncloud.ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Broker{
			Replicas: 2,
			Resources: &sncloud.ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1DefaultNodeResource{
				Cpu:    "1000m",
				Memory: "4294967296",
			},
		},
		DisplayName:    ptr.To(clusterName),
		InstanceName:   instanceName,
		Location:       *selectedLocation,
		ReleaseChannel: ptr.To("rapid"),
	}

	instJSON, err := json.Marshal(inst)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal instance: %v", err)
	}
	clusJSON, err := json.Marshal(clus)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal cluster: %v", err)
	}

	messages := []mcp.PromptMessage{
		{
			Content: mcp.TextContent{
				Type: "text",
				Text: "The following is the Pulsar instance JSON definition and the Pulsar cluster JSON definition, you can use the `sncloud_resources_apply` tool to apply the resources to the StreamNative Cloud. Please directly use the JSON content and not modify the content. The PulsarCluster name is required to be empty. You will need to apply PulsarInstance first, then apply PulsarCluster.",
			},
			Role: mcp.RoleUser,
		},
		{
			Content: mcp.TextContent{
				Type: "text",
				Text: string(instJSON),
			},
			Role: mcp.RoleUser,
		},
		{
			Content: mcp.TextContent{
				Type: "text",
				Text: string(clusJSON),
			},
			Role: mcp.RoleUser,
		},
	}

	return &mcp.GetPromptResult{
		Description: fmt.Sprintf("Create a new Serverless Pulsar cluster %s's related resources that can be applied to the StreamNative Cloud.", clusterName),
		Messages:    messages,
	}, nil
}

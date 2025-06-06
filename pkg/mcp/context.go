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

var McpContextPulsarInstance *string
var McpContextPulsarCluster *string
var McpContextOrganization *string

func SetMcpContext(instance, cluster, organization string) {
	McpContextPulsarInstance = &instance
	McpContextPulsarCluster = &cluster
	McpContextOrganization = &organization
}

// GetMcpContext returns the currently configured StreamNative Cloud context.
// If the context has not been fully configured, it returns empty strings.
func GetMcpContext() (string, string, string) {
	if McpContextPulsarInstance == nil || McpContextPulsarCluster == nil || McpContextOrganization == nil {
		return "", "", ""
	}

	return *McpContextPulsarInstance, *McpContextPulsarCluster, *McpContextOrganization
}

func ResetMcpContext() {
	McpContextPulsarInstance = nil
	McpContextPulsarCluster = nil
	McpContextOrganization = nil
}

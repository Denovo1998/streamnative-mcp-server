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

/*
Api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sncloud

import (
	"encoding/json"
)

// ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig ExecConfig specifies a command to provide client credentials. The command is exec'd and outputs structured stdout holding credentials.  See the client.authentiction.k8s.io API group for specifications of the exact input and output format
type ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig struct {
	// Preferred input version of the ExecInfo. The returned ExecCredentials MUST use the same encoding version as the input.
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Arguments to pass to the command when executing it.
	Args []string `json:"args,omitempty"`
	// Command to execute.
	Command string `json:"command"`
	// Env defines additional environment variables to expose to the process. These are unioned with the host's environment, as well as variables client-go uses to pass argument to the plugin.
	Env []ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecEnvVar `json:"env,omitempty"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig(command string) *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig{}
	this.Command = command
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfigWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfigWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) GetArgs() []string {
	if o == nil || o.Args == nil {
		var ret []string
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) GetArgsOk() ([]string, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) HasArgs() bool {
	if o != nil && o.Args != nil {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the Args field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) SetArgs(v []string) {
	o.Args = v
}

// GetCommand returns the Command field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) SetCommand(v string) {
	o.Command = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) GetEnv() []ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecEnvVar {
	if o == nil || o.Env == nil {
		var ret []ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecEnvVar
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) GetEnvOk() ([]ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecEnvVar, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) HasEnv() bool {
	if o != nil && o.Env != nil {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecEnvVar and assigns it to the Env field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) SetEnv(v []ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecEnvVar) {
	o.Env = v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if true {
		toSerialize["command"] = o.Command
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) Get() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) Set(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ExecConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



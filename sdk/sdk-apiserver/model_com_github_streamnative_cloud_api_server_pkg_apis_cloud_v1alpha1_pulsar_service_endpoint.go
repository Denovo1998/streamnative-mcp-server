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

// ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint struct for ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint
type ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint struct {
	DnsName string `json:"dnsName"`
	// Gateway is the name of the PulsarGateway to use for the endpoint, will be empty if endpointAccess is not configured.
	Gateway *string `json:"gateway,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint(dnsName string) *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint{}
	this.DnsName = dnsName
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpointWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpointWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint{}
	return &this
}

// GetDnsName returns the DnsName field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) GetDnsName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) GetDnsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnsName, true
}

// SetDnsName sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) SetDnsName(v string) {
	o.DnsName = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) SetGateway(v string) {
	o.Gateway = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) SetType(v string) {
	o.Type = &v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnsName"] = o.DnsName
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) Get() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) Set(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarServiceEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



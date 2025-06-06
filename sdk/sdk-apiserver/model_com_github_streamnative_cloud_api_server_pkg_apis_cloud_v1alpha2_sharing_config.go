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

// ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig struct for ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig
type ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig struct {
	Namespaces []string `json:"namespaces"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig(namespaces []string) *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig{}
	this.Namespaces = namespaces
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfigWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfigWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig{}
	return &this
}

// GetNamespaces returns the Namespaces field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig) GetNamespaces() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig) GetNamespacesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespaces, true
}

// SetNamespaces sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig) SetNamespaces(v []string) {
	o.Namespaces = v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["namespaces"] = o.Namespaces
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig) Get() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig) Set(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig) *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha2SharingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



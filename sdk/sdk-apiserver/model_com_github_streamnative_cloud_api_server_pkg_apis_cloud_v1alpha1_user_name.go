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

// ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName struct for ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName
type ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName struct {
	First string `json:"first"`
	Last string `json:"last"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName(first string, last string) *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName{}
	this.First = first
	this.Last = last
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserNameWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserNameWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName{}
	return &this
}

// GetFirst returns the First field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) GetFirst() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.First
}

// GetFirstOk returns a tuple with the First field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) GetFirstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.First, true
}

// SetFirst sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) SetFirst(v string) {
	o.First = v
}

// GetLast returns the Last field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) GetLast() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Last
}

// GetLastOk returns a tuple with the Last field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) GetLastOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Last, true
}

// SetLast sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) SetLast(v string) {
	o.Last = v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["first"] = o.First
	}
	if true {
		toSerialize["last"] = o.Last
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) Get() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) Set(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1UserName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



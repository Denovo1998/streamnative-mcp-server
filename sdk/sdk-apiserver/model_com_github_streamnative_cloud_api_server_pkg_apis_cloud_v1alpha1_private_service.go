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

// ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService struct for ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService
type ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService struct {
	// AllowedIds is the list of Ids that are allowed to connect to the private endpoint service, only can be configured when the access type is private, private endpoint service will be disabled if the whitelist is empty.
	AllowedIds []string `json:"allowedIds,omitempty"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService{}
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateServiceWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateServiceWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService{}
	return &this
}

// GetAllowedIds returns the AllowedIds field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) GetAllowedIds() []string {
	if o == nil || o.AllowedIds == nil {
		var ret []string
		return ret
	}
	return o.AllowedIds
}

// GetAllowedIdsOk returns a tuple with the AllowedIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) GetAllowedIdsOk() ([]string, bool) {
	if o == nil || o.AllowedIds == nil {
		return nil, false
	}
	return o.AllowedIds, true
}

// HasAllowedIds returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) HasAllowedIds() bool {
	if o != nil && o.AllowedIds != nil {
		return true
	}

	return false
}

// SetAllowedIds gets a reference to the given []string and assigns it to the AllowedIds field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) SetAllowedIds(v []string) {
	o.AllowedIds = v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedIds != nil {
		toSerialize["allowedIds"] = o.AllowedIds
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) Get() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) Set(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



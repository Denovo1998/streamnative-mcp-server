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

// ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus SelfSubjectRbacReviewStatus defines the observed state of SelfSubjectRulesReview
type ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus struct {
	UserPrivileges *string `json:"userPrivileges,omitempty"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus instantiates a new ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus() *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus {
	this := ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus{}
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatusWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatusWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus {
	this := ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus{}
	return &this
}

// GetUserPrivileges returns the UserPrivileges field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus) GetUserPrivileges() string {
	if o == nil || o.UserPrivileges == nil {
		var ret string
		return ret
	}
	return *o.UserPrivileges
}

// GetUserPrivilegesOk returns a tuple with the UserPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus) GetUserPrivilegesOk() (*string, bool) {
	if o == nil || o.UserPrivileges == nil {
		return nil, false
	}
	return o.UserPrivileges, true
}

// HasUserPrivileges returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus) HasUserPrivileges() bool {
	if o != nil && o.UserPrivileges != nil {
		return true
	}

	return false
}

// SetUserPrivileges gets a reference to the given string and assigns it to the UserPrivileges field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus) SetUserPrivileges(v string) {
	o.UserPrivileges = &v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserPrivileges != nil {
		toSerialize["userPrivileges"] = o.UserPrivileges
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus) Get() *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus) Set(val *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus(val *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus) *NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SelfSubjectRbacReviewStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



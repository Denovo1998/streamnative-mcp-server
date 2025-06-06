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

// ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec IdentityPoolSpec defines the desired state of IdentityPool
type ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec struct {
	AuthType string `json:"authType"`
	Description string `json:"description"`
	Expression string `json:"expression"`
	ProviderName string `json:"providerName"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec(authType string, description string, expression string, providerName string) *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec{}
	this.AuthType = authType
	this.Description = description
	this.Expression = expression
	this.ProviderName = providerName
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpecWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpecWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec{}
	return &this
}

// GetAuthType returns the AuthType field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) GetAuthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) SetAuthType(v string) {
	o.AuthType = v
}

// GetDescription returns the Description field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) SetDescription(v string) {
	o.Description = v
}

// GetExpression returns the Expression field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) SetExpression(v string) {
	o.Expression = v
}

// GetProviderName returns the ProviderName field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) SetProviderName(v string) {
	o.ProviderName = v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authType"] = o.AuthType
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["expression"] = o.Expression
	}
	if true {
		toSerialize["providerName"] = o.ProviderName
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) Get() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) Set(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1IdentityPoolSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls struct for ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls
type ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls struct {
	// The TLS secret to use (in the namespace of the gateway workload pods).
	CertSecretName string `json:"certSecretName"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls(certSecretName string) *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls{}
	this.CertSecretName = certSecretName
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTlsWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTlsWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls{}
	return &this
}

// GetCertSecretName returns the CertSecretName field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls) GetCertSecretName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertSecretName
}

// GetCertSecretNameOk returns a tuple with the CertSecretName field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls) GetCertSecretNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertSecretName, true
}

// SetCertSecretName sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls) SetCertSecretName(v string) {
	o.CertSecretName = v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["certSecretName"] = o.CertSecretName
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls) Get() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls) Set(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls) *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberIstioGatewayTls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



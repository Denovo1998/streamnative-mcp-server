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

// ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec struct for ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec
type ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec struct {
	// Access is the access type of the pulsar gateway, available values are public or private. It is immutable, with the default value public.
	Access *string `json:"access,omitempty"`
	// Domains is the list of domain suffix that the pulsar gateway will serve. This is automatically generated based on the PulsarGateway name and PoolMember domain.
	Domains []ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Domain `json:"domains,omitempty"`
	PoolMemberRef *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberReference `json:"poolMemberRef,omitempty"`
	PrivateService *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService `json:"privateService,omitempty"`
	// TopologyAware is the configuration of the topology aware feature of the pulsar gateway.
	TopologyAware map[string]interface{} `json:"topologyAware,omitempty"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec{}
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpecWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpecWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) GetAccess() string {
	if o == nil || o.Access == nil {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) GetAccessOk() (*string, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) SetAccess(v string) {
	o.Access = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) GetDomains() []ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Domain {
	if o == nil || o.Domains == nil {
		var ret []ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Domain
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) GetDomainsOk() ([]ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Domain, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Domain and assigns it to the Domains field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) SetDomains(v []ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Domain) {
	o.Domains = v
}

// GetPoolMemberRef returns the PoolMemberRef field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) GetPoolMemberRef() ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberReference {
	if o == nil || o.PoolMemberRef == nil {
		var ret ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberReference
		return ret
	}
	return *o.PoolMemberRef
}

// GetPoolMemberRefOk returns a tuple with the PoolMemberRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) GetPoolMemberRefOk() (*ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberReference, bool) {
	if o == nil || o.PoolMemberRef == nil {
		return nil, false
	}
	return o.PoolMemberRef, true
}

// HasPoolMemberRef returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) HasPoolMemberRef() bool {
	if o != nil && o.PoolMemberRef != nil {
		return true
	}

	return false
}

// SetPoolMemberRef gets a reference to the given ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberReference and assigns it to the PoolMemberRef field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) SetPoolMemberRef(v ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PoolMemberReference) {
	o.PoolMemberRef = &v
}

// GetPrivateService returns the PrivateService field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) GetPrivateService() ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService {
	if o == nil || o.PrivateService == nil {
		var ret ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService
		return ret
	}
	return *o.PrivateService
}

// GetPrivateServiceOk returns a tuple with the PrivateService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) GetPrivateServiceOk() (*ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService, bool) {
	if o == nil || o.PrivateService == nil {
		return nil, false
	}
	return o.PrivateService, true
}

// HasPrivateService returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) HasPrivateService() bool {
	if o != nil && o.PrivateService != nil {
		return true
	}

	return false
}

// SetPrivateService gets a reference to the given ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService and assigns it to the PrivateService field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) SetPrivateService(v ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PrivateService) {
	o.PrivateService = &v
}

// GetTopologyAware returns the TopologyAware field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) GetTopologyAware() map[string]interface{} {
	if o == nil || o.TopologyAware == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TopologyAware
}

// GetTopologyAwareOk returns a tuple with the TopologyAware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) GetTopologyAwareOk() (map[string]interface{}, bool) {
	if o == nil || o.TopologyAware == nil {
		return nil, false
	}
	return o.TopologyAware, true
}

// HasTopologyAware returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) HasTopologyAware() bool {
	if o != nil && o.TopologyAware != nil {
		return true
	}

	return false
}

// SetTopologyAware gets a reference to the given map[string]interface{} and assigns it to the TopologyAware field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) SetTopologyAware(v map[string]interface{}) {
	o.TopologyAware = v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	if o.PoolMemberRef != nil {
		toSerialize["poolMemberRef"] = o.PoolMemberRef
	}
	if o.PrivateService != nil {
		toSerialize["privateService"] = o.PrivateService
	}
	if o.TopologyAware != nil {
		toSerialize["topologyAware"] = o.TopologyAware
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) Get() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) Set(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1PulsarGatewaySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



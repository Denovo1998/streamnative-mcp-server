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

// ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy AutoScalingPolicy defines the min/max replicas for component is autoscaling enabled.
type ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy struct {
	MaxReplicas int32 `json:"maxReplicas"`
	MinReplicas *int32 `json:"minReplicas,omitempty"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy(maxReplicas int32) *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy{}
	this.MaxReplicas = maxReplicas
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicyWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicyWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy{}
	return &this
}

// GetMaxReplicas returns the MaxReplicas field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) GetMaxReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxReplicas
}

// GetMaxReplicasOk returns a tuple with the MaxReplicas field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) GetMaxReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxReplicas, true
}

// SetMaxReplicas sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) SetMaxReplicas(v int32) {
	o.MaxReplicas = v
}

// GetMinReplicas returns the MinReplicas field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) GetMinReplicas() int32 {
	if o == nil || o.MinReplicas == nil {
		var ret int32
		return ret
	}
	return *o.MinReplicas
}

// GetMinReplicasOk returns a tuple with the MinReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) GetMinReplicasOk() (*int32, bool) {
	if o == nil || o.MinReplicas == nil {
		return nil, false
	}
	return o.MinReplicas, true
}

// HasMinReplicas returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) HasMinReplicas() bool {
	if o != nil && o.MinReplicas != nil {
		return true
	}

	return false
}

// SetMinReplicas gets a reference to the given int32 and assigns it to the MinReplicas field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) SetMinReplicas(v int32) {
	o.MinReplicas = &v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["maxReplicas"] = o.MaxReplicas
	}
	if o.MinReplicas != nil {
		toSerialize["minReplicas"] = o.MinReplicas
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) Get() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) Set(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1AutoScalingPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



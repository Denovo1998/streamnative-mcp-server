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
	"time"
)

// ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus VvpDeploymentRunningStatus defines the observed state of VvpDeployment
type ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus struct {
	// Conditions is an array of current observed conditions.
	Conditions []ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentStatusCondition `json:"conditions,omitempty"`
	JobId *string `json:"jobId,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	TransitionTime *time.Time `json:"transitionTime,omitempty"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus instantiates a new ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus() *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus {
	this := ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus{}
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatusWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatusWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus {
	this := ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) GetConditions() []ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentStatusCondition {
	if o == nil || o.Conditions == nil {
		var ret []ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentStatusCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) GetConditionsOk() ([]ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentStatusCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentStatusCondition and assigns it to the Conditions field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) SetConditions(v []ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentStatusCondition) {
	o.Conditions = v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) GetJobId() string {
	if o == nil || o.JobId == nil {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) GetJobIdOk() (*string, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) SetJobId(v string) {
	o.JobId = &v
}

// GetTransitionTime returns the TransitionTime field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) GetTransitionTime() time.Time {
	if o == nil || o.TransitionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.TransitionTime
}

// GetTransitionTimeOk returns a tuple with the TransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) GetTransitionTimeOk() (*time.Time, bool) {
	if o == nil || o.TransitionTime == nil {
		return nil, false
	}
	return o.TransitionTime, true
}

// HasTransitionTime returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) HasTransitionTime() bool {
	if o != nil && o.TransitionTime != nil {
		return true
	}

	return false
}

// SetTransitionTime gets a reference to the given time.Time and assigns it to the TransitionTime field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) SetTransitionTime(v time.Time) {
	o.TransitionTime = &v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.JobId != nil {
		toSerialize["jobId"] = o.JobId
	}
	if o.TransitionTime != nil {
		toSerialize["transitionTime"] = o.TransitionTime
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) Get() *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) Set(val *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus(val *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) *NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentRunningStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



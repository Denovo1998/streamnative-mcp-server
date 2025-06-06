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

// ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition Condition represents an observation of an object's state. Conditions are an extension mechanism intended to be used when the details of an observation are not a priori known or would not apply to all instances of a given Kind.  Conditions should be added to explicitly convey properties that users and components care about rather than requiring those properties to be inferred from other observations. Once defined, the meaning of a Condition can not be changed arbitrarily - it becomes part of the API, and has the same backwards- and forwards-compatibility concerns of any other part of the API.
type ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	Message *string `json:"message,omitempty"`
	// observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Status string `json:"status"`
	Type string `json:"type"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition(status string, type_ string) *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition{}
	this.Status = status
	this.Type = type_
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ConditionWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1ConditionWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition {
	this := ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition{}
	return &this
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) GetLastTransitionTime() time.Time {
	if o == nil || o.LastTransitionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) GetLastTransitionTimeOk() (*time.Time, bool) {
	if o == nil || o.LastTransitionTime == nil {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) HasLastTransitionTime() bool {
	if o != nil && o.LastTransitionTime != nil {
		return true
	}

	return false
}

// SetLastTransitionTime gets a reference to the given time.Time and assigns it to the LastTransitionTime field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) SetLastTransitionTime(v time.Time) {
	o.LastTransitionTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) SetMessage(v string) {
	o.Message = &v
}

// GetObservedGeneration returns the ObservedGeneration field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) GetObservedGeneration() int64 {
	if o == nil || o.ObservedGeneration == nil {
		var ret int64
		return ret
	}
	return *o.ObservedGeneration
}

// GetObservedGenerationOk returns a tuple with the ObservedGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) GetObservedGenerationOk() (*int64, bool) {
	if o == nil || o.ObservedGeneration == nil {
		return nil, false
	}
	return o.ObservedGeneration, true
}

// HasObservedGeneration returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) HasObservedGeneration() bool {
	if o != nil && o.ObservedGeneration != nil {
		return true
	}

	return false
}

// SetObservedGeneration gets a reference to the given int64 and assigns it to the ObservedGeneration field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) SetObservedGeneration(v int64) {
	o.ObservedGeneration = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) SetType(v string) {
	o.Type = v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastTransitionTime != nil {
		toSerialize["lastTransitionTime"] = o.LastTransitionTime
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ObservedGeneration != nil {
		toSerialize["observedGeneration"] = o.ObservedGeneration
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) Get() *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) Set(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition(val *ComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisCloudV1alpha1Condition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



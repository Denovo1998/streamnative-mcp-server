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

// V1Condition Condition contains details for one aspect of the current state of this API Resource.
type V1Condition struct {
	// lastTransitionTime is the last time the condition transitioned from one status to another. This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.
	LastTransitionTime time.Time `json:"lastTransitionTime"`
	// message is a human readable message indicating details about the transition. This may be an empty string.
	Message string `json:"message"`
	// observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	// reason contains a programmatic identifier indicating the reason for the condition's last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty.
	Reason string `json:"reason"`
	// status of the condition, one of True, False, Unknown.
	Status string `json:"status"`
	// type of condition in CamelCase or in foo.example.com/CamelCase.
	Type string `json:"type"`
}

// NewV1Condition instantiates a new V1Condition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Condition(lastTransitionTime time.Time, message string, reason string, status string, type_ string) *V1Condition {
	this := V1Condition{}
	this.LastTransitionTime = lastTransitionTime
	this.Message = message
	this.Reason = reason
	this.Status = status
	this.Type = type_
	return &this
}

// NewV1ConditionWithDefaults instantiates a new V1Condition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ConditionWithDefaults() *V1Condition {
	this := V1Condition{}
	return &this
}

// GetLastTransitionTime returns the LastTransitionTime field value
func (o *V1Condition) GetLastTransitionTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value
// and a boolean to check if the value has been set.
func (o *V1Condition) GetLastTransitionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastTransitionTime, true
}

// SetLastTransitionTime sets field value
func (o *V1Condition) SetLastTransitionTime(v time.Time) {
	o.LastTransitionTime = v
}

// GetMessage returns the Message field value
func (o *V1Condition) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *V1Condition) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *V1Condition) SetMessage(v string) {
	o.Message = v
}

// GetObservedGeneration returns the ObservedGeneration field value if set, zero value otherwise.
func (o *V1Condition) GetObservedGeneration() int64 {
	if o == nil || o.ObservedGeneration == nil {
		var ret int64
		return ret
	}
	return *o.ObservedGeneration
}

// GetObservedGenerationOk returns a tuple with the ObservedGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Condition) GetObservedGenerationOk() (*int64, bool) {
	if o == nil || o.ObservedGeneration == nil {
		return nil, false
	}
	return o.ObservedGeneration, true
}

// HasObservedGeneration returns a boolean if a field has been set.
func (o *V1Condition) HasObservedGeneration() bool {
	if o != nil && o.ObservedGeneration != nil {
		return true
	}

	return false
}

// SetObservedGeneration gets a reference to the given int64 and assigns it to the ObservedGeneration field.
func (o *V1Condition) SetObservedGeneration(v int64) {
	o.ObservedGeneration = &v
}

// GetReason returns the Reason field value
func (o *V1Condition) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *V1Condition) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *V1Condition) SetReason(v string) {
	o.Reason = v
}

// GetStatus returns the Status field value
func (o *V1Condition) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *V1Condition) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *V1Condition) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *V1Condition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V1Condition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V1Condition) SetType(v string) {
	o.Type = v
}

func (o V1Condition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["lastTransitionTime"] = o.LastTransitionTime
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.ObservedGeneration != nil {
		toSerialize["observedGeneration"] = o.ObservedGeneration
	}
	if true {
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

type NullableV1Condition struct {
	value *V1Condition
	isSet bool
}

func (v NullableV1Condition) Get() *V1Condition {
	return v.value
}

func (v *NullableV1Condition) Set(val *V1Condition) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Condition) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Condition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Condition(val *V1Condition) *NullableV1Condition {
	return &NullableV1Condition{value: val, isSet: true}
}

func (v NullableV1Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Condition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



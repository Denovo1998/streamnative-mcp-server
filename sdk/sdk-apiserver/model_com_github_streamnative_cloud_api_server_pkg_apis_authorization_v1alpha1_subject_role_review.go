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

// ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview SubjectRoleReview
type ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`
	Spec *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewSpec `json:"spec,omitempty"`
	Status *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewStatus `json:"status,omitempty"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview instantiates a new ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview() *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview {
	this := ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview{}
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview {
	this := ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) GetMetadata() V1ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret V1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) GetMetadataOk() (*V1ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ObjectMeta and assigns it to the Metadata field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) SetMetadata(v V1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) GetSpec() ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewSpec {
	if o == nil || o.Spec == nil {
		var ret ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) GetSpecOk() (*ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewSpec and assigns it to the Spec field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) SetSpec(v ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) GetStatus() ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewStatus {
	if o == nil || o.Status == nil {
		var ret ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) GetStatusOk() (*ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewStatus and assigns it to the Status field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) SetStatus(v ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReviewStatus) {
	o.Status = &v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) Get() *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) Set(val *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview(val *ComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) *NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisAuthorizationV1alpha1SubjectRoleReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



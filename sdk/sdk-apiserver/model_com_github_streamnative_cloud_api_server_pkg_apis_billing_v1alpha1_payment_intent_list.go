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

// ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList struct for ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList
type ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	Items []ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntent `json:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *V1ListMeta `json:"metadata,omitempty"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList instantiates a new ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList(items []ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntent) *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList {
	this := ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList{}
	this.Items = items
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentListWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentListWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList {
	this := ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) GetItems() []ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntent {
	if o == nil {
		var ret []ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntent
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) GetItemsOk() ([]ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) SetItems(v []ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntent) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) GetMetadata() V1ListMeta {
	if o == nil || o.Metadata == nil {
		var ret V1ListMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) GetMetadataOk() (*V1ListMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given V1ListMeta and assigns it to the Metadata field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) SetMetadata(v V1ListMeta) {
	o.Metadata = &v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if true {
		toSerialize["items"] = o.Items
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) Get() *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) Set(val *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList(val *ComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) *NullableComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisBillingV1alpha1PaymentIntentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



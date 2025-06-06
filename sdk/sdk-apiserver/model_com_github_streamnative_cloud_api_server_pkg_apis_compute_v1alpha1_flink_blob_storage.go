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

// ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage FlinkBlobStorage defines the configuration for the Flink blob storage.
type ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage struct {
	// Bucket is required if you want to use cloud storage.
	Bucket *string `json:"bucket,omitempty"`
	// Path is the sub path in the bucket. Leave it empty if you want to use the whole bucket.
	Path *string `json:"path,omitempty"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage instantiates a new ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage() *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage {
	this := ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage{}
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorageWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorageWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage {
	this := ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) SetBucket(v string) {
	o.Bucket = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) SetPath(v string) {
	o.Path = &v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) Get() *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) Set(val *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage(val *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) *NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1FlinkBlobStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



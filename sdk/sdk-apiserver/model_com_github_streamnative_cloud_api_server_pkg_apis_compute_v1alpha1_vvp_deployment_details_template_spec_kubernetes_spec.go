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

// ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec VvpDeploymentDetailsTemplateSpecKubernetesSpec defines the desired state of VvpDeploymentDetails
type ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec struct {
	JobManagerPodTemplate *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1PodTemplate `json:"jobManagerPodTemplate,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
	TaskManagerPodTemplate *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1PodTemplate `json:"taskManagerPodTemplate,omitempty"`
}

// NewComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec instantiates a new ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec() *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec {
	this := ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec{}
	return &this
}

// NewComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpecWithDefaults instantiates a new ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpecWithDefaults() *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec {
	this := ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec{}
	return &this
}

// GetJobManagerPodTemplate returns the JobManagerPodTemplate field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) GetJobManagerPodTemplate() ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1PodTemplate {
	if o == nil || o.JobManagerPodTemplate == nil {
		var ret ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1PodTemplate
		return ret
	}
	return *o.JobManagerPodTemplate
}

// GetJobManagerPodTemplateOk returns a tuple with the JobManagerPodTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) GetJobManagerPodTemplateOk() (*ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1PodTemplate, bool) {
	if o == nil || o.JobManagerPodTemplate == nil {
		return nil, false
	}
	return o.JobManagerPodTemplate, true
}

// HasJobManagerPodTemplate returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) HasJobManagerPodTemplate() bool {
	if o != nil && o.JobManagerPodTemplate != nil {
		return true
	}

	return false
}

// SetJobManagerPodTemplate gets a reference to the given ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1PodTemplate and assigns it to the JobManagerPodTemplate field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) SetJobManagerPodTemplate(v ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1PodTemplate) {
	o.JobManagerPodTemplate = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetTaskManagerPodTemplate returns the TaskManagerPodTemplate field value if set, zero value otherwise.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) GetTaskManagerPodTemplate() ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1PodTemplate {
	if o == nil || o.TaskManagerPodTemplate == nil {
		var ret ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1PodTemplate
		return ret
	}
	return *o.TaskManagerPodTemplate
}

// GetTaskManagerPodTemplateOk returns a tuple with the TaskManagerPodTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) GetTaskManagerPodTemplateOk() (*ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1PodTemplate, bool) {
	if o == nil || o.TaskManagerPodTemplate == nil {
		return nil, false
	}
	return o.TaskManagerPodTemplate, true
}

// HasTaskManagerPodTemplate returns a boolean if a field has been set.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) HasTaskManagerPodTemplate() bool {
	if o != nil && o.TaskManagerPodTemplate != nil {
		return true
	}

	return false
}

// SetTaskManagerPodTemplate gets a reference to the given ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1PodTemplate and assigns it to the TaskManagerPodTemplate field.
func (o *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) SetTaskManagerPodTemplate(v ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1PodTemplate) {
	o.TaskManagerPodTemplate = &v
}

func (o ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JobManagerPodTemplate != nil {
		toSerialize["jobManagerPodTemplate"] = o.JobManagerPodTemplate
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.TaskManagerPodTemplate != nil {
		toSerialize["taskManagerPodTemplate"] = o.TaskManagerPodTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec struct {
	value *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec
	isSet bool
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) Get() *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec {
	return v.value
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) Set(val *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec(val *ComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) *NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec {
	return &NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec{value: val, isSet: true}
}

func (v NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComGithubStreamnativeCloudApiServerPkgApisComputeV1alpha1VvpDeploymentDetailsTemplateSpecKubernetesSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type GlueJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueJobSpec   `json:"spec,omitempty"`
	Status            GlueJobStatus `json:"status,omitempty"`
}

type GlueJobSpecCommand struct {
	// +optional
	Name           string `json:"name,omitempty" tf:"name,omitempty"`
	ScriptLocation string `json:"scriptLocation" tf:"script_location"`
}

type GlueJobSpecExecutionProperty struct {
	// +optional
	MaxConcurrentRuns int64 `json:"maxConcurrentRuns,omitempty" tf:"max_concurrent_runs,omitempty"`
}

type GlueJobSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// Deprecated
	AllocatedCapacity int64 `json:"allocatedCapacity,omitempty" tf:"allocated_capacity,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Command []GlueJobSpecCommand `json:"command" tf:"command"`
	// +optional
	Connections []string `json:"connections,omitempty" tf:"connections,omitempty"`
	// +optional
	DefaultArguments map[string]string `json:"defaultArguments,omitempty" tf:"default_arguments,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ExecutionProperty []GlueJobSpecExecutionProperty `json:"executionProperty,omitempty" tf:"execution_property,omitempty"`
	// +optional
	MaxCapacity float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`
	// +optional
	MaxRetries int64  `json:"maxRetries,omitempty" tf:"max_retries,omitempty"`
	Name       string `json:"name" tf:"name"`
	RoleArn    string `json:"roleArn" tf:"role_arn"`
	// +optional
	SecurityConfiguration string `json:"securityConfiguration,omitempty" tf:"security_configuration,omitempty"`
	// +optional
	Timeout int64 `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type GlueJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GlueJobSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlueJobList is a list of GlueJobs
type GlueJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlueJob CRD objects
	Items []GlueJob `json:"items,omitempty"`
}

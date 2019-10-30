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

type AthenaDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AthenaDatabaseSpec   `json:"spec,omitempty"`
	Status            AthenaDatabaseStatus `json:"status,omitempty"`
}

type AthenaDatabaseSpecEncryptionConfiguration struct {
	EncryptionOption string `json:"encryptionOption" tf:"encryption_option"`
	// +optional
	KmsKey string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`
}

type AthenaDatabaseSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Bucket string `json:"bucket" tf:"bucket"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionConfiguration []AthenaDatabaseSpecEncryptionConfiguration `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`
	// +optional
	ForceDestroy bool   `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`
	Name         string `json:"name" tf:"name"`
}

type AthenaDatabaseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AthenaDatabaseSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AthenaDatabaseList is a list of AthenaDatabases
type AthenaDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AthenaDatabase CRD objects
	Items []AthenaDatabase `json:"items,omitempty"`
}

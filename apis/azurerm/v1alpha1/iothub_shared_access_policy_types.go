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

type IothubSharedAccessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubSharedAccessPolicySpec   `json:"spec,omitempty"`
	Status            IothubSharedAccessPolicyStatus `json:"status,omitempty"`
}

type IothubSharedAccessPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	DeviceConnect bool   `json:"deviceConnect,omitempty" tf:"device_connect,omitempty"`
	IothubName    string `json:"iothubName" tf:"iothub_name"`
	Name          string `json:"name" tf:"name"`
	// +optional
	PrimaryConnectionString string `json:"-" sensitive:"true" tf:"primary_connection_string,omitempty"`
	// +optional
	PrimaryKey string `json:"-" sensitive:"true" tf:"primary_key,omitempty"`
	// +optional
	RegistryRead bool `json:"registryRead,omitempty" tf:"registry_read,omitempty"`
	// +optional
	RegistryWrite     bool   `json:"registryWrite,omitempty" tf:"registry_write,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryConnectionString string `json:"-" sensitive:"true" tf:"secondary_connection_string,omitempty"`
	// +optional
	SecondaryKey string `json:"-" sensitive:"true" tf:"secondary_key,omitempty"`
	// +optional
	ServiceConnect bool `json:"serviceConnect,omitempty" tf:"service_connect,omitempty"`
}

type IothubSharedAccessPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IothubSharedAccessPolicySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IothubSharedAccessPolicyList is a list of IothubSharedAccessPolicys
type IothubSharedAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IothubSharedAccessPolicy CRD objects
	Items []IothubSharedAccessPolicy `json:"items,omitempty"`
}

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

type VirtualNetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkSpec   `json:"spec,omitempty"`
	Status            VirtualNetworkStatus `json:"status,omitempty"`
}

type VirtualNetworkSpecDdosProtectionPlan struct {
	Enable bool   `json:"enable" tf:"enable"`
	ID     string `json:"ID" tf:"id"`
}

type VirtualNetworkSpecSubnet struct {
	AddressPrefix string `json:"addressPrefix" tf:"address_prefix"`
	// +optional
	ID   string `json:"ID,omitempty" tf:"id,omitempty"`
	Name string `json:"name" tf:"name"`
	// +optional
	SecurityGroup string `json:"securityGroup,omitempty" tf:"security_group,omitempty"`
}

type VirtualNetworkSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MinItems=1
	AddressSpace []string `json:"addressSpace" tf:"address_space"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DdosProtectionPlan []VirtualNetworkSpecDdosProtectionPlan `json:"ddosProtectionPlan,omitempty" tf:"ddos_protection_plan,omitempty"`
	// +optional
	DnsServers        []string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`
	Location          string   `json:"location" tf:"location"`
	Name              string   `json:"name" tf:"name"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Subnet []VirtualNetworkSpecSubnet `json:"subnet,omitempty" tf:"subnet,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VirtualNetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VirtualNetworkSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualNetworkList is a list of VirtualNetworks
type VirtualNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualNetwork CRD objects
	Items []VirtualNetwork `json:"items,omitempty"`
}

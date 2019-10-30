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

type NetworkSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityGroupSpec   `json:"spec,omitempty"`
	Status            NetworkSecurityGroupStatus `json:"status,omitempty"`
}

type NetworkSecurityGroupSpecSecurityRule struct {
	Access string `json:"access" tf:"access"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	DestinationAddressPrefix string `json:"destinationAddressPrefix,omitempty" tf:"destination_address_prefix,omitempty"`
	// +optional
	DestinationAddressPrefixes []string `json:"destinationAddressPrefixes,omitempty" tf:"destination_address_prefixes,omitempty"`
	// +optional
	DestinationApplicationSecurityGroupIDS []string `json:"destinationApplicationSecurityGroupIDS,omitempty" tf:"destination_application_security_group_ids,omitempty"`
	// +optional
	DestinationPortRange string `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`
	// +optional
	DestinationPortRanges []string `json:"destinationPortRanges,omitempty" tf:"destination_port_ranges,omitempty"`
	Direction             string   `json:"direction" tf:"direction"`
	Name                  string   `json:"name" tf:"name"`
	Priority              int      `json:"priority" tf:"priority"`
	Protocol              string   `json:"protocol" tf:"protocol"`
	// +optional
	SourceAddressPrefix string `json:"sourceAddressPrefix,omitempty" tf:"source_address_prefix,omitempty"`
	// +optional
	SourceAddressPrefixes []string `json:"sourceAddressPrefixes,omitempty" tf:"source_address_prefixes,omitempty"`
	// +optional
	SourceApplicationSecurityGroupIDS []string `json:"sourceApplicationSecurityGroupIDS,omitempty" tf:"source_application_security_group_ids,omitempty"`
	// +optional
	SourcePortRange string `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`
	// +optional
	SourcePortRanges []string `json:"sourcePortRanges,omitempty" tf:"source_port_ranges,omitempty"`
}

type NetworkSecurityGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecurityRule []NetworkSecurityGroupSpecSecurityRule `json:"securityRule,omitempty" tf:"security_rule,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NetworkSecurityGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NetworkSecurityGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkSecurityGroupList is a list of NetworkSecurityGroups
type NetworkSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkSecurityGroup CRD objects
	Items []NetworkSecurityGroup `json:"items,omitempty"`
}

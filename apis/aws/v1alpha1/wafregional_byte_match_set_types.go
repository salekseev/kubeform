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

type WafregionalByteMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalByteMatchSetSpec   `json:"spec,omitempty"`
	Status            WafregionalByteMatchSetStatus `json:"status,omitempty"`
}

type WafregionalByteMatchSetSpecByteMatchTuplesFieldToMatch struct {
	// +optional
	Data string `json:"data,omitempty" tf:"data,omitempty"`
	Type string `json:"type" tf:"type"`
}

type WafregionalByteMatchSetSpecByteMatchTuples struct {
	// +kubebuilder:validation:MaxItems=1
	FieldToMatch         []WafregionalByteMatchSetSpecByteMatchTuplesFieldToMatch `json:"fieldToMatch" tf:"field_to_match"`
	PositionalConstraint string                                                   `json:"positionalConstraint" tf:"positional_constraint"`
	// +optional
	TargetString       string `json:"targetString,omitempty" tf:"target_string,omitempty"`
	TextTransformation string `json:"textTransformation" tf:"text_transformation"`
}

type WafregionalByteMatchSetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ByteMatchTuples []WafregionalByteMatchSetSpecByteMatchTuples `json:"byteMatchTuples,omitempty" tf:"byte_match_tuples,omitempty"`
	Name            string                                       `json:"name" tf:"name"`
}

type WafregionalByteMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *WafregionalByteMatchSetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WafregionalByteMatchSetList is a list of WafregionalByteMatchSets
type WafregionalByteMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WafregionalByteMatchSet CRD objects
	Items []WafregionalByteMatchSet `json:"items,omitempty"`
}

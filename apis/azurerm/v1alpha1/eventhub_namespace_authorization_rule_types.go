package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type EventhubNamespaceAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubNamespaceAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            EventhubNamespaceAuthorizationRuleStatus `json:"status,omitempty"`
}

type EventhubNamespaceAuthorizationRuleSpec struct {
	// +optional
	Listen bool `json:"listen,omitempty"`
	// +optional
	// Deprecated
	Location string `json:"location,omitempty"`
	// +optional
	Manage            bool   `json:"manage,omitempty"`
	Name              string `json:"name"`
	NamespaceName     string `json:"namespace_name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	Send bool `json:"send,omitempty"`
}

type EventhubNamespaceAuthorizationRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventhubNamespaceAuthorizationRuleList is a list of EventhubNamespaceAuthorizationRules
type EventhubNamespaceAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventhubNamespaceAuthorizationRule CRD objects
	Items []EventhubNamespaceAuthorizationRule `json:"items,omitempty"`
}
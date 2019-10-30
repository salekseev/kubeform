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

type IamAccountPasswordPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamAccountPasswordPolicySpec   `json:"spec,omitempty"`
	Status            IamAccountPasswordPolicyStatus `json:"status,omitempty"`
}

type IamAccountPasswordPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllowUsersToChangePassword bool `json:"allowUsersToChangePassword,omitempty" tf:"allow_users_to_change_password,omitempty"`
	// +optional
	ExpirePasswords bool `json:"expirePasswords,omitempty" tf:"expire_passwords,omitempty"`
	// +optional
	HardExpiry bool `json:"hardExpiry,omitempty" tf:"hard_expiry,omitempty"`
	// +optional
	MaxPasswordAge int `json:"maxPasswordAge,omitempty" tf:"max_password_age,omitempty"`
	// +optional
	MinimumPasswordLength int `json:"minimumPasswordLength,omitempty" tf:"minimum_password_length,omitempty"`
	// +optional
	PasswordReusePrevention int `json:"passwordReusePrevention,omitempty" tf:"password_reuse_prevention,omitempty"`
	// +optional
	RequireLowercaseCharacters bool `json:"requireLowercaseCharacters,omitempty" tf:"require_lowercase_characters,omitempty"`
	// +optional
	RequireNumbers bool `json:"requireNumbers,omitempty" tf:"require_numbers,omitempty"`
	// +optional
	RequireSymbols bool `json:"requireSymbols,omitempty" tf:"require_symbols,omitempty"`
	// +optional
	RequireUppercaseCharacters bool `json:"requireUppercaseCharacters,omitempty" tf:"require_uppercase_characters,omitempty"`
}

type IamAccountPasswordPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IamAccountPasswordPolicySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamAccountPasswordPolicyList is a list of IamAccountPasswordPolicys
type IamAccountPasswordPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamAccountPasswordPolicy CRD objects
	Items []IamAccountPasswordPolicy `json:"items,omitempty"`
}

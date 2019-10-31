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

type GoogleServiceAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleServiceAccountSpec   `json:"spec,omitempty"`
	Status            GoogleServiceAccountStatus `json:"status,omitempty"`
}

type GoogleServiceAccountSpec struct {
	// +optional
	SecretRef   *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	Source string `json:"source" tf:"source"`

	// +optional
	// If assigning billing role, specificy a billing account (default is to assign at the organizational level).
	BillingAccountID string `json:"billingAccountID,omitempty" tf:"billing_account_id,omitempty"`
	// +optional
	// Generate keys for service accounts.
	GenerateKeys bool `json:"generateKeys,omitempty" tf:"generate_keys,omitempty"`
	// +optional
	// Grant billing user role.
	GrantBillingRole bool `json:"grantBillingRole,omitempty" tf:"grant_billing_role,omitempty"`
	// +optional
	// Grant roles for shared VPC management.
	GrantXpnRoles bool `json:"grantXpnRoles,omitempty" tf:"grant_xpn_roles,omitempty"`
	// +optional
	// Names of the service accounts to create.
	Names []string `json:"names,omitempty" tf:"names,omitempty"`
	// +optional
	// Id of the organization for org-level roles.
	OrgID string `json:"orgID,omitempty" tf:"org_id,omitempty"`
	// +optional
	// Prefix applied to service account names.
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
	// +optional
	// Project id where service account will be created.
	ProjectID string `json:"projectID,omitempty" tf:"project_id,omitempty"`
	// +optional
	// Common roles to apply to all service accounts, project=>role as elements.
	ProjectRoles []string `json:"projectRoles,omitempty" tf:"project_roles,omitempty"`
}

type GoogleServiceAccountOutput struct {
	// Service account email (for single use).
	// +optional
	Email string `json:"email" tf:"email"`
	// Service account emails.
	// +optional
	Emails string `json:"emails" tf:"emails"`
	// Service account emails.
	// +optional
	EmailsList string `json:"emailsList" tf:"emails_list"`
	// IAM-format service account email (for single use).
	// +optional
	IamEmail string `json:"iamEmail" tf:"iam_email"`
	// IAM-format service account emails.
	// +optional
	IamEmails string `json:"iamEmails" tf:"iam_emails"`
	// IAM-format service account emails.
	// +optional
	IamEmailsList string `json:"iamEmailsList" tf:"iam_emails_list"`
	// Service account key (for single use).
	// +optional
	Key string `json:"key" tf:"key"`
	// Map of service account keys.
	// +optional
	Keys string `json:"keys" tf:"keys"`
	// Service account resource (for single use).
	// +optional
	ServiceAccount string `json:"serviceAccount" tf:"service_account"`
	// Service account resources.
	// +optional
	ServiceAccounts string `json:"serviceAccounts" tf:"service_accounts"`
}

type GoogleServiceAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GoogleServiceAccountOutput `json:"output,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleServiceAccountList is a list of GoogleServiceAccounts
type GoogleServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleServiceAccount CRD objects
	Items []GoogleServiceAccount `json:"items,omitempty"`
}
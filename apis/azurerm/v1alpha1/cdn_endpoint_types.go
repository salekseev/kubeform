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

type CdnEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CdnEndpointSpec   `json:"spec,omitempty"`
	Status            CdnEndpointStatus `json:"status,omitempty"`
}

type CdnEndpointSpecGeoFilter struct {
	Action       string   `json:"action" tf:"action"`
	CountryCodes []string `json:"countryCodes" tf:"country_codes"`
	RelativePath string   `json:"relativePath" tf:"relative_path"`
}

type CdnEndpointSpecOrigin struct {
	HostName string `json:"hostName" tf:"host_name"`
	// +optional
	HttpPort int64 `json:"httpPort,omitempty" tf:"http_port,omitempty"`
	// +optional
	HttpsPort int64  `json:"httpsPort,omitempty" tf:"https_port,omitempty"`
	Name      string `json:"name" tf:"name"`
}

type CdnEndpointSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ContentTypesToCompress []string `json:"contentTypesToCompress,omitempty" tf:"content_types_to_compress,omitempty"`
	// +optional
	GeoFilter []CdnEndpointSpecGeoFilter `json:"geoFilter,omitempty" tf:"geo_filter,omitempty"`
	// +optional
	HostName string `json:"hostName,omitempty" tf:"host_name,omitempty"`
	// +optional
	IsCompressionEnabled bool `json:"isCompressionEnabled,omitempty" tf:"is_compression_enabled,omitempty"`
	// +optional
	IsHTTPAllowed bool `json:"isHTTPAllowed,omitempty" tf:"is_http_allowed,omitempty"`
	// +optional
	IsHTTPSAllowed bool   `json:"isHTTPSAllowed,omitempty" tf:"is_https_allowed,omitempty"`
	Location       string `json:"location" tf:"location"`
	Name           string `json:"name" tf:"name"`
	// +optional
	OptimizationType string                  `json:"optimizationType,omitempty" tf:"optimization_type,omitempty"`
	Origin           []CdnEndpointSpecOrigin `json:"origin" tf:"origin"`
	// +optional
	OriginHostHeader string `json:"originHostHeader,omitempty" tf:"origin_host_header,omitempty"`
	// +optional
	OriginPath string `json:"originPath,omitempty" tf:"origin_path,omitempty"`
	// +optional
	ProbePath   string `json:"probePath,omitempty" tf:"probe_path,omitempty"`
	ProfileName string `json:"profileName" tf:"profile_name"`
	// +optional
	QuerystringCachingBehaviour string `json:"querystringCachingBehaviour,omitempty" tf:"querystring_caching_behaviour,omitempty"`
	ResourceGroupName           string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CdnEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CdnEndpointSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CdnEndpointList is a list of CdnEndpoints
type CdnEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CdnEndpoint CRD objects
	Items []CdnEndpoint `json:"items,omitempty"`
}

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

type LambdaLayerVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaLayerVersionSpec   `json:"spec,omitempty"`
	Status            LambdaLayerVersionStatus `json:"status,omitempty"`
}

type LambdaLayerVersionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	CompatibleRuntimes []string `json:"compatibleRuntimes,omitempty" tf:"compatible_runtimes,omitempty"`
	// +optional
	CreatedDate string `json:"createdDate,omitempty" tf:"created_date,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Filename string `json:"filename,omitempty" tf:"filename,omitempty"`
	// +optional
	LayerArn  string `json:"layerArn,omitempty" tf:"layer_arn,omitempty"`
	LayerName string `json:"layerName" tf:"layer_name"`
	// +optional
	LicenseInfo string `json:"licenseInfo,omitempty" tf:"license_info,omitempty"`
	// +optional
	S3Bucket string `json:"s3Bucket,omitempty" tf:"s3_bucket,omitempty"`
	// +optional
	S3Key string `json:"s3Key,omitempty" tf:"s3_key,omitempty"`
	// +optional
	S3ObjectVersion string `json:"s3ObjectVersion,omitempty" tf:"s3_object_version,omitempty"`
	// +optional
	SourceCodeHash string `json:"sourceCodeHash,omitempty" tf:"source_code_hash,omitempty"`
	// +optional
	SourceCodeSize int `json:"sourceCodeSize,omitempty" tf:"source_code_size,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
}

type LambdaLayerVersionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LambdaLayerVersionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LambdaLayerVersionList is a list of LambdaLayerVersions
type LambdaLayerVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LambdaLayerVersion CRD objects
	Items []LambdaLayerVersion `json:"items,omitempty"`
}

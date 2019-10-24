package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VirtualEnvSpec defines the desired state of VirtualEnv
// +k8s:openapi-gen=true
type VirtualEnvSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Header to keep env name in trace
	Header string `json:"header,omitempty"`
	// Environment variable to mark env name of deployment
	EvMark string `json:"evmark,omitempty"`
}

// VirtualEnvStatus defines the observed state of VirtualEnv
// +k8s:openapi-gen=true
type VirtualEnvStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Available env names
	EnvNames []string `json:"envs"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualEnv is the Schema for the virtualenvs API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=virtualenvs,scope=Namespaced
type VirtualEnv struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualEnvSpec   `json:"spec,omitempty"`
	Status VirtualEnvStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualEnvList contains a list of VirtualEnv
type VirtualEnvList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualEnv `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualEnv{}, &VirtualEnvList{})
}
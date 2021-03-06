package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ComputerSpec defines the desired state of Computer
// +k8s:openapi-gen=true
type ComputerSpec struct {
	// ID of the computer
	ID int `json:"id,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Computer is the Schema for the computers API
// +genclient
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=computers,scope=Namespaced
type Computer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputerSpec      `json:"spec,omitempty"`
	Status corev1.NodeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputerList contains a list of Computer
type ComputerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Computer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Computer{}, &ComputerList{})
}

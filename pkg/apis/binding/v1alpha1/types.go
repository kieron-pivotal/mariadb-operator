package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MysqlBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []MysqlBinding `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MysqlBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              MysqlBindingSpec   `json:"spec"`
	Status            MysqlBindingStatus `json:"status,omitempty"`
}

type MysqlBindingSpec struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Hostname string `json:"hostname"`
}
type MysqlBindingStatus struct {
	Status string `json:"status" default:"pending"`
}

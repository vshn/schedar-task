package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true

// SimplePostgreSQL is the API for creating a Stackgres PostgreSQL instance
type SimplePostgreSQL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SimplePostgreSQLSpec   `json:"spec"`
	Status SimplePostgreSQLStatus `json:"status,omitempty"`
}

// SimplePostgreSQLSpec defines the desired state of a SimplePostgreSQL.
type SimplePostgreSQLSpec struct {
	//TODO implement me
}

// SimplePostgreSQLStatus defines the desired state of a SimplePostgreSQL.
type SimplePostgreSQLStatus struct{}

// +kubebuilder:object:root=true

// SimplePostgreSQLList defines a list of SimplePostgreSQL
type SimplePostgreSQLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []SimplePostgreSQL `json:"items"`
}

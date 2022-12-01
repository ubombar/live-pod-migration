package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LivePodMigration represents a live migration of a Pod
type LivePodMigration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LivePodMigrationSpec   `json:"spec"`
	Status LivePodMigrationStatus `json:"status"`
}

// LivePodMigration is the spec for a LivePodMigration resource
type LivePodMigrationSpec struct {
	PodName     string `json:"podName"`
	ServiceName string `json:"serviceName"`
}

// LivePodMigrationStatus is the status for a LivePodMigration resource
type LivePodMigrationStatus struct {
	MigrationStatus  string `json:"migrationStatus"`
	MigrationMessage string `json:"migrationMessage"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LivePodMigrationList is a list of LivePodMigration resources
type LivePodMigrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []LivePodMigration `json:"items"`
}

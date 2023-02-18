package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LivePodMigrationRequest represents a live migration of a Pod
type LivePodMigrationRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LivePodMigrationRequestSpec   `json:"spec"`
	Status LivePodMigrationRequestStatus `json:"status"`
}

type LivePodMigrationRequestSpec struct {
	PodName             string          `json:"podName"`
	SourceNodeName      string          `json:"sourceNodeName"`
	DestinationNodeName string          `json:"destinationNodeName"`
	Method              MigrationMethod `json:"method"`
}

type MigrationMethod string

const (
	ColdMigration MigrationMethod = "ColdMigration"
)

type LivePodMigrationRequestStatus struct {
}

// // LivePodMigration is the spec for a LivePodMigration resource
// type LivePodMigrationSpec struct {
// 	PodNamespace    string `json:"podNamespace"`
// 	PodName         string `json:"podName"`
// 	DestinationNode string `json:"destinationNode"`
// 	ServiceName     string `json:"serviceName"`
// }

// // LivePodMigrationStatus is the status for a LivePodMigration resource
// type LivePodMigrationStatus struct {
// 	MigrationStatus  MigrationStatus `json:"migrationStatus"`
// 	MigrationMessage string          `json:"migrationMessage"`
// 	CheckpointFile   string          `json:"checkpointFile"`
// 	PodAccessible    bool            `json:"podAccessible"`
// }

// type MigrationStatus string

// const (
// 	MigrationStatusPending       MigrationStatus = "Pending"
// 	MigrationStatusError         MigrationStatus = "Error"
// 	MigrationStatusCheckpointing MigrationStatus = "Checkpointing"
// 	MigrationStatusTransferring  MigrationStatus = "Transferring"
// 	MigrationStatusRestoring     MigrationStatus = "Restoring"
// 	MigrationStatusCleaning      MigrationStatus = "Cleaning"
// 	MigrationStatusCompleted     MigrationStatus = "Completed"
// )

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LivePodMigrationList is a list of LivePodMigration resources
type LivePodMigrationRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []LivePodMigrationRequest `json:"items"`
}

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
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
	MethodColdMigration MigrationMethod = "ColdMigration"
)

type LivePodMigrationRequestStatus struct {
	// Descritption regarding to the migration request. It might be empty if no errors.
	Description string `json:"description"`

	// Describes the state of the request.
	RequestState LivePodMigrationRequestState `json:"requestState"`
}

type LivePodMigrationRequestState string

const (
	LivePodMigrationRequestStatePending    LivePodMigrationRequestState = "Pending"
	LivePodMigrationRequestStateInProgress LivePodMigrationRequestState = "InProgress"
	LivePodMigrationRequestStateFailed     LivePodMigrationRequestState = "Failed"
	LivePodMigrationRequestStateSucceed    LivePodMigrationRequestState = "Succeed"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LivePodMigrationList is a list of LivePodMigration resources
type LivePodMigrationRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []LivePodMigrationRequest `json:"items"`
}

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
	PodNamespace    string `json:"podNamespace"`
	PodName         string `json:"podName"`
	DestinationNode string `json:"destinationNode"`
	ServiceName     string `json:"serviceName"`
}

// LivePodMigrationStatus is the status for a LivePodMigration resource
type LivePodMigrationStatus struct {
	MigrationStatus  MigrationStatus `json:"migrationStatus"`
	MigrationMessage string          `json:"migrationMessage"`
	CheckpointFile   string          `json:"checkpointFile"`
	PodAccessible    bool            `json:"podAccessible"`
}

type MigrationStatus string

const (
	MigrationStatusPending       MigrationStatus = "Pending"
	MigrationStatusError         MigrationStatus = "Error"
	MigrationStatusCheckpointing MigrationStatus = "Checkpointing"
	MigrationStatusTransferring  MigrationStatus = "Transferring"
	MigrationStatusRestoring     MigrationStatus = "Restoring"
	MigrationStatusCleaning      MigrationStatus = "Cleaning"
	MigrationStatusCompleted     MigrationStatus = "Completed"
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
type LivePodMigrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []LivePodMigration `json:"items"`
}

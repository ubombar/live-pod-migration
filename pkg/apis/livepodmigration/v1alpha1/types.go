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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ColdMigration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ColdMigrationSpec   `json:"spec"`
	Status ColdMigrationStatus `json:"status"`
}

type ColdMigrationSpec struct {
	LivePodMigrationRequestName string `json:"livePodMigrationRequestName"`
}

type ColdMigrationStatus struct {
	// Current stage of the migration
	CurrentMigrationStage string `json:"migrationStage"`

	// This can be InProgress | Finnished | Failed
	SourceSyncer           ColdMigrationSyncer `json:"sourceSyncer"`
	SourceDescriptions     string              `json:"sourceDescriptions"`
	DestinationSyncer      ColdMigrationSyncer `json:"destinationSyncer"`
	DestinationDescription string              `json:"destinationDescription"`
}

type ColdMigrationSyncer string

const (
	ColdMigrationSyncerInProgress  ColdMigrationSyncer = "InProgress"
	ColdMigrationSyncerInFinnished ColdMigrationSyncer = "Finnished"
	ColdMigrationSyncerInFailed    ColdMigrationSyncer = "Failed"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LivePodMigrationList is a list of LivePodMigration resources
type ColdMigrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ColdMigration `json:"items"`
}

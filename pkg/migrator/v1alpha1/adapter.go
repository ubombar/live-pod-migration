package v1alpha1

import "github.com/ubombar/live-pod-migration/pkg/apis/livepodmigration/v1alpha1"

const (
	StatusPending       = "StatusPending"
	StatusCheckpointing = "StatusCheckpointing"
	StatusTransfering   = "StatusTransfering"
	StatusRestoring     = "StatusRestoring"
	StatusDone          = "StatusDone"
	StatusError         = "StatusError"
)

type MigratorAdapterConfig struct {
	// IPv4 address of self and peer migrator
	SelfAddress string
	PeerAddress string

	// Status of self and peer
	SelfStatus string
	PeerStatus string

	// Denotes if the current instance acts as a client
	IsClient bool

	LPMInfo *v1alpha1.LivePodMigration
}

// This is the migrator adapter. A stateful adapter interface for client and server
// migrator communication as well as kubernetes.
type MigrationAdapter interface {
	// Initialize the variables for implementation.
	Configure(mc *MigratorAdapterConfig) bool

	// Establish connection
	DialPeer() bool

	// Abort communication with reason and description
	Abort(reason, description string) string

	// Get self and peer states
	PeerState() string
	SelfState() string

	// In case there is a difference in known states, this syncs the states of self and peer
	SyncStates(selfState string) string

	// From this point we have migration related functions.

	InitiateMigration() bool

	// Blocking function. Reads and transmits the checkpoint file.
	TransitCheckpoint() bool
}

type MigrationAdapterRPC struct {
	MigrationAdapter
}

func NewMigrationAdapterRPC(mc *MigratorAdapterConfig) *MigrationAdapterRPC {
	adapter := &MigrationAdapterRPC{}

	adapter.Configure(mc)

	return adapter
}

// Initialize the variables for implementation.
func (a *MigrationAdapterRPC) Configure(mc *MigratorAdapterConfig) bool {
	return false
}

// Establish connection
func (a *MigrationAdapterRPC) DialPeer() bool {
	return false
}

// Abort communication with reason and description
func (a *MigrationAdapterRPC) Abort(reason, description string) string {
	return ""
}

// Get self and peer states
func (a *MigrationAdapterRPC) PeerState() string {
	return ""
}

func (a *MigrationAdapterRPC) SelfState() string {
	return ""
}

// In case there is a difference in known states, this syncs the states of self and peer
func (a *MigrationAdapterRPC) SyncStates(selfState string) string {
	return ""
}

// From this point we have migration related functions.
func (a *MigrationAdapterRPC) InitiateMigration() bool {
	return false
}

// Blocking function. Reads and transmits the checkpoint file.
func (a *MigrationAdapterRPC) TransitCheckpoint() bool {
	return false
}

package v1alpha1

import "github.com/ubombar/live-pod-migration/pkg/apis/livepodmigration/v1alpha1"

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

func NewMigrationAdapterRPC() *MigrationAdapterRPC {
	return &MigrationAdapterRPC{}
}

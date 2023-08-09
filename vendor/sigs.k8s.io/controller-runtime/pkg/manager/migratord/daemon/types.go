package daemon

import (
	"fmt"
	"strings"
	"time"
)

const (
	IncomingConsumer      = "incoming-consumer"
	PreparingConsumer     = "preparing-consumer"
	CheckpointingConsumer = "checkpointing-consumer"
	TransferingConsumer   = "transfering-consumer"
	RestoringConsumer     = "restoring-consumer"
	DoneConsumer          = "done-consumer"
	ErrorConsumer         = "error-consumer"
	SyncConsumer          = "sync-consumer"
)

const (
	IncomingQueue      = "incoming-queue"
	PreparingQueue     = "preparing-queue"
	CheckpointingQueue = "checkpointing-queue"
	TransferingQueue   = "transfering-queue"
	RestoringQueue     = "restoring-queue"
	DoneQueue          = "done-queue"
	ErrorQueue         = "error-queue"
	SyncQueue          = "sync-queue"
	NullQueue          = ""
)

const (
	MigrationJobStore  = "migration-job-store"
	MigrationRoleStore = "migration-role-store"
	MigrationSyncStore = "migration-sync-store"
)

type MigrationRole string

const (
	MigrationRoleServer MigrationRole = "migration-role-server"
	MigrationRoleClient MigrationRole = "migration-role-client"
)

func (r MigrationRole) PeersRole() MigrationRole {
	if r == MigrationRoleClient {
		return MigrationRoleServer
	} else {
		return MigrationRoleClient
	}
}

type MigrationStatus string

func (m MigrationStatus) ToQueueName() string {
	return fmt.Sprintf("%v-queue", strings.Split(string(m), "-")[1])
}

const (
	StatusIncoming      MigrationStatus = "status-incoming"
	StatusPreparing     MigrationStatus = "status-preparing"
	StatusCheckpointing MigrationStatus = "status-checkpointing"
	StatusTransfering   MigrationStatus = "status-transfering"
	StatusRestoring     MigrationStatus = "status-restoring"
	StatusDone          MigrationStatus = "status-done"
	StatusError         MigrationStatus = "status-error"
	StatusNull          MigrationStatus = "status-null"
)

type MigrationMethod string

const (
	MethodBasic    MigrationMethod = "method-basic"
	MethodPrecopy  MigrationMethod = "method-precopy"
	MethodPostcopy MigrationMethod = "method-postcopy"
)

type DaemonConfig struct {
	SelfAddress         string
	SelfPort            int
	QueueSize           int
	CheckpointDirectory string
}

func DefaultDaemonConfigs() DaemonConfig {
	return DaemonConfig{
		SelfAddress:         "localhost",
		SelfPort:            9213,
		QueueSize:           64,
		CheckpointDirectory: DefaultCheckpointDirectory,
	}
}

func DefaultDaemonConfigsWithAddress(ipaddress string) DaemonConfig {
	return DaemonConfig{
		SelfAddress:         ipaddress,
		SelfPort:            9213,
		QueueSize:           64,
		CheckpointDirectory: DefaultCheckpointDirectory,
	}
}

const DefaultCheckpointDirectory = "/var/lib/migratord/checkpoints"

type MigrationJob struct {
	// Used for describing the migration
	MigrationId string

	// Client and Servers ip address
	ClientIP string
	ServerIP string

	// Client and Servers port
	ClientPort int
	ServerPort int

	// Cotnainer identification number
	ClientContainerID string
	ServerContainerID string

	// Migration status describes in which stage migration is
	Status MigrationStatus

	// This represents the failure reason
	Error error

	// Shows if the container is running
	Running bool

	// Private key
	PrivateKey string

	// Server's username
	Username string

	// Creation date
	CreationDate time.Time

	// Role of the migratord that owns this object
	Role MigrationRole

	// How the migration will be performed
	Method MigrationMethod
}

func (j MigrationJob) AddressString(role MigrationRole) string {
	if role == MigrationRoleClient {
		return fmt.Sprintf("%v:%d", j.ClientIP, j.ClientPort)
	} else {
		return fmt.Sprintf("%v:%d", j.ServerIP, j.ServerPort)
	}
}

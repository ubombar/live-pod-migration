package daemon

const (
	IncomingConsumer      = "incoming-consumer"
	PreparingConsumer     = "preparing-consumer"
	CheckpointingConsumer = "checkpointing-consumer"
	TransferingConsumer   = "transfering-consumer"
	RestoringConsumer     = "restoring-consumer"
)

const (
	IncomingQueue      = "incoming-queue"
	PreparingQueue     = "preparing-queue"
	CheckpointingQueue = "checkpointing-queue"
	TransferingQueue   = "transfering-queue"
	RestoringQueue     = "restoring-queue"
)

const (
	MigrationJobStore  = "migration-job-store"
	MigrationRoleStore = "migration-role-store"
)

type MigrationRole string

const (
	MigrationRoleServer MigrationRole = "migration-role-server"
	MigrationRoleClient MigrationRole = "migration-role-client"
)

type DaemonConfig struct {
	SelfAddress string
	SelfPort    int
	QueueSize   int
}

func DefaultDaemonConfigs() DaemonConfig {
	return DaemonConfig{
		SelfAddress: "localhost",
		SelfPort:    4545,
		QueueSize:   64,
	}
}

func DefaultDaemonConfigsWithAddress(ipaddress string) DaemonConfig {
	return DaemonConfig{
		SelfAddress: ipaddress,
		SelfPort:    4545,
		QueueSize:   64,
	}
}

type MigrationJob struct {
	// Lots of things here
}

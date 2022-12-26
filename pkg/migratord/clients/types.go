package clients

const (
	ClientDocker = "docker"
)

type ContainerInfo struct {
	ContainerId    string
	ImageId        string
	ContainerNames []string

	// Other stuff for storage and networking
}

type CheckpointInfo struct {
	ContainerId   string
	CheckpointDir string

	// Checkpoint id is the name of the checkpoint.
	CheckpointId string
}

type Client interface {
	// Get runtime eg 'docker'
	Runtime() string

	// Get version of the runtime
	Version() string

	// Checkpointing functions
	GetContainerInfo(containerId string) (*ContainerInfo, error)
	GetCheckpointInfo(checkpointid string) (*CheckpointInfo, error)
}

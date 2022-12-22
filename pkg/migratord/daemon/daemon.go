package daemon

import (
	"github.com/sirupsen/logrus"
	"github.com/ubombar/live-pod-migration/pkg/migratord/clients"
	"github.com/ubombar/live-pod-migration/pkg/migratord/consumers"
	"github.com/ubombar/live-pod-migration/pkg/migratord/structures"
)

type Daemon interface {
	// Start and stop
	Start() error
	Stop() error

	// Access interfaces
	GetConsumer(name string) consumers.Consumer
	GetQueue(name string) structures.Queue
	GetJobStore() structures.Store
	GetSyncStore() structures.Store
	GetRoleStore() structures.Store
	GetContainerClient(name string) clients.Client
	GetDefaultContainerClient() clients.Client
}

type daemon struct {
	config    *DaemonConfig
	consumers map[string]consumers.Consumer
	queues    map[string]structures.Queue
	jobstore  structures.Store
	syncstore structures.Store
	rolestore structures.Store
	client    map[string]clients.Client
}

func NewDaemon(config *DaemonConfig) *daemon {
	d := &daemon{
		config: config,
	}

	// Set the queues
	d.queues = map[string]structures.Queue{
		IncomingQueue:      structures.NewQueue(IncomingQueue, config.QueueSize),
		PreparingQueue:     structures.NewQueue(PreparingQueue, config.QueueSize),
		CheckpointingQueue: structures.NewQueue(CheckpointingQueue, config.QueueSize),
		TransferingQueue:   structures.NewQueue(TransferingQueue, config.QueueSize),
		RestoringQueue:     structures.NewQueue(RestoringQueue, config.QueueSize),
	}

	// Set the store
	d.jobstore = structures.NewStore(MigrationJobStore)
	d.rolestore = structures.NewStore(MigrationRoleStore)
	d.syncstore = structures.NewStore(MigrationSyncStore)

	// Set the consumers
	d.consumers = map[string]consumers.Consumer{
		IncomingConsumer:      consumers.NewConsumer(d.GetQueue(IncomingQueue), d.incomingCallback),
		PreparingConsumer:     consumers.NewConsumer(d.GetQueue(PreparingQueue), d.preparingCallback),
		CheckpointingConsumer: consumers.NewConsumer(d.GetQueue(CheckpointingQueue), d.checkpointingCallback),
		TransferingConsumer:   consumers.NewConsumer(d.GetQueue(TransferingQueue), d.transferingCallback),
		RestoringConsumer:     consumers.NewConsumer(d.GetQueue(RestoringQueue), d.restoringCallback),
	}

	// Set client
	d.client = map[string]clients.Client{
		clients.ClientDocker: clients.NewDockerClient(),
	}

	return d
}

func (d *daemon) Start() error {
	// Start running all of the consumers
	for _, consumer := range d.consumers {
		consumer.Run()
	}

	// Start gRPC interface
	return nil
}

func (d *daemon) Stop() error {
	return nil
}

func (d *daemon) GetConsumer(name string) consumers.Consumer {
	if con, ok := d.consumers[name]; ok {
		return con
	}

	logrus.Errorln("cannot find specified consumer:", name)
	return nil
}

func (d *daemon) GetQueue(name string) structures.Queue {
	if que, ok := d.queues[name]; ok {
		return que
	}

	logrus.Errorln("cannot find specified queue:", name)
	return nil
}

func (d *daemon) GetJobStore() structures.Store {
	return d.jobstore
}

func (d *daemon) GetSyncStore() structures.Store {
	return d.syncstore
}

func (d *daemon) GetRoleStore() structures.Store {
	return d.rolestore
}

func (d *daemon) GetContainerClient(name string) clients.Client {
	if cli, ok := d.client[name]; ok {
		return cli
	}

	logrus.Errorln("cannot find specified client:", name)
	return nil
}

func (d *daemon) GetDefaultContainerClient() clients.Client {
	return d.GetContainerClient(clients.ClientDocker)
}

// Client: Checks if the container and runtime exists, server is reachable. Also
// shares the job with the server.
//
// Server: Checks if the container, runtime and image exists. If good, changes the
// job's status to preparing and notifies the client.
func (d *daemon) incomingCallback(id string) error {
	// TODO: Requires implementation
	// job, _ := d.GetJobStore().Fetch(id)
	// role, _ := d.GetRoleStore().Fetch(id)

	return nil
}

// Handle migration
func (d *daemon) preparingCallback(id string) error {
	// TODO: Requires implementation
	return nil
}

// Handle migration
func (d *daemon) checkpointingCallback(id string) error {
	// TODO: Requires implementation
	return nil
}

// Handle migration
func (d *daemon) transferingCallback(id string) error {
	// TODO: Requires implementation
	return nil
}

// Handle migration
func (d *daemon) restoringCallback(id string) error {
	// TODO: Requires implementation
	return nil
}

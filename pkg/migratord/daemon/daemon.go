package daemon

import (
	"errors"

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
	GetRoleStore() structures.Store
	GetSyncer() Syncer
	GetContainerClient(name string) clients.Client
	GetDefaultContainerClient() clients.Client
	GetRPC() RPC

	// Access configuration
	GetConfig() DaemonConfig
}

type daemon struct {
	config    *DaemonConfig
	consumers map[string]consumers.Consumer
	queues    map[string]structures.Queue
	jobstore  structures.Store
	rolestore structures.Store

	syncer Syncer
	client map[string]clients.Client
	grpc   RPC
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
		DoneQueue:          structures.NewQueue(DoneQueue, config.QueueSize),
		ErrorQueue:         structures.NewQueue(ErrorQueue, config.QueueSize),
		SyncQueue:          structures.NewQueue(SyncQueue, config.QueueSize),
	}

	// Set the store
	d.jobstore = structures.NewStore(MigrationJobStore)
	d.rolestore = structures.NewStore(MigrationRoleStore)

	// Set the consumers
	d.consumers = map[string]consumers.Consumer{
		IncomingConsumer:      consumers.NewConsumer(d.GetQueue(IncomingQueue), d.incomingCallback),
		PreparingConsumer:     consumers.NewConsumer(d.GetQueue(PreparingQueue), d.preparingCallback),
		CheckpointingConsumer: consumers.NewConsumer(d.GetQueue(CheckpointingQueue), d.checkpointingCallback),
		TransferingConsumer:   consumers.NewConsumer(d.GetQueue(TransferingQueue), d.transferingCallback),
		RestoringConsumer:     consumers.NewConsumer(d.GetQueue(RestoringQueue), d.restoringCallback),
		DoneConsumer:          consumers.NewConsumer(d.GetQueue(DoneQueue), d.doneCallback),
		ErrorConsumer:         consumers.NewConsumer(d.GetQueue(ErrorQueue), d.errorCallback),
		SyncConsumer:          consumers.NewConsumer(d.GetQueue(SyncQueue), d.syncCallback),
	}

	d.syncer = NewSyncer(d)

	// Set grpc
	d.grpc = NewRPC(RPCConfig{
		Address: config.SelfAddress,
		Port:    config.SelfPort,
	}, d)

	// Set client
	d.client = map[string]clients.Client{
		clients.ClientPodman: clients.NewPodmanClient(),
	}

	return d
}

func (d *daemon) Start() error {
	// Start running all of the consumers
	for _, consumer := range d.consumers {
		if err := consumer.Run(); err != nil {
			logrus.Errorln("cannot run consumer")
			return err
		}
	}

	// Start gRPC interface
	err := d.grpc.Run()

	if err != nil {
		logrus.Errorln("cannot run grpc")
		return err
	}

	logrus.Infoln("Started queue-consumers and rpc interface.")

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

func (d *daemon) GetRoleStore() structures.Store {
	return d.rolestore
}

func (d *daemon) GetSyncer() Syncer {
	return d.syncer
}

func (d *daemon) GetContainerClient(name string) clients.Client {
	if cli, ok := d.client[name]; ok {
		return cli
	}

	logrus.Errorln("cannot find specified client:", name)
	return nil
}

func (d *daemon) GetDefaultContainerClient() clients.Client {
	return d.GetContainerClient(clients.ClientPodman)
}

func (d *daemon) GetRPC() RPC {
	return d.grpc
}

func (d *daemon) GetConfig() DaemonConfig {
	return *d.config
}

// func (d *daemon) getMigrationObjects(migrationid string) (*MigrationJob, *MigrationRole, error) {
// 	obj, err := d.GetRoleStore().Fetch(migrationid)

// 	if err != nil {
// 		return nil, nil, errors.New("cannot get role of migrationid")
// 	}

// 	role, ok := obj.(MigrationRole)

// 	if !ok {
// 		return nil, nil, errors.New("migration store does not contain role")
// 	}

// 	obj, err = d.GetJobStore().Fetch(migrationid)

// 	if err != nil {
// 		return nil, nil, errors.New("cannot get job of migrationid")
// 	}

// 	job, ok := obj.(MigrationJob)

// 	if !ok {
// 		return nil, nil, errors.New("migration store does not contain job")
// 	}

// 	return &job, &role, nil
// }

func FromMigrationId(migrationid string, daemon Daemon) (*MigrationJob, *MigrationRole, *Sync, error) {
	obj, err := daemon.GetSyncer().GetSyncStore().Fetch(migrationid)

	if err != nil {
		return nil, nil, nil, errors.New("cannot get sync of migrationid")
	}

	sync, ok := obj.(Sync)

	if !ok {
		return nil, nil, nil, errors.New("sync store does not contain sync")
	}

	obj, err = daemon.GetRoleStore().Fetch(migrationid)

	if err != nil {
		return nil, nil, nil, errors.New("cannot get role of migrationid")
	}

	role, ok := obj.(MigrationRole)

	if !ok {
		return nil, nil, nil, errors.New("migration store does not contain role")
	}

	obj, err = daemon.GetJobStore().Fetch(migrationid)

	if err != nil {
		return nil, nil, nil, errors.New("cannot get job of migrationid")
	}

	job, ok := obj.(MigrationJob)

	if !ok {
		return nil, nil, nil, errors.New("migration store does not contain job")
	}

	return &job, &role, &sync, nil
}

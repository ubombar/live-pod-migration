package daemon

import (
	"errors"

	"github.com/ubombar/live-pod-migration/pkg/migratord/structures"
)

// Denotes a sync object. This is used for waiting the other peer until both
// reach sync points.
type Sync struct {
	// Have client finnished pricessing for current status
	ClientDone bool

	// Have server finnished pricessing for current status
	ServerDone bool

	// Current and next status for job
	CurrentStatus MigrationStatus
	NextStatus    MigrationStatus

	// Name of the queue it will be appended after finishing this state
	NextQueueName string

	// Migration id
	MigrationId string
}

// By the role, marks the done segment true
func (s *Sync) SetCompleted(role MigrationRole) {
	if role == MigrationRoleServer {
		s.ServerDone = true
	} else {
		s.ClientDone = true
	}
}

// If both lights are green, then it means current stage has finnished.
func (s *Sync) AllCompleted() bool {
	return s.ClientDone && s.ServerDone
}

// Return if the current role is completed
func (s *Sync) RoleCompleted(role MigrationRole) bool {
	if role == MigrationRoleServer {
		return s.ServerDone
	} else {
		return s.ClientDone
	}
}

type Syncer interface {
	// Register the migration if with requested attributes. current status and the next job queue.
	RegisterJob(migrationid string, currentStatus MigrationStatus, nextQueueName string)

	// Indicates the job has finnished. If the peer still processing it will not add the migration
	// to the next queue until peer finishes.
	FinishJob(migrationid string, role MigrationRole) error

	// Get the stores
	GetSyncStore() structures.Store
}

type syncer struct {
	Syncer
	d         Daemon
	syncstore structures.Store
}

func NewSyncer(daemon Daemon) *syncer {
	s := &syncer{
		d:         daemon,
		syncstore: structures.NewStore(MigrationSyncStore),
	}

	return s
}

func (s *syncer) RegisterJob(migrationid string, currentStatus MigrationStatus, nextQueueName string) {
	sync := Sync{
		ClientDone:    false,
		ServerDone:    false,
		CurrentStatus: currentStatus,
		MigrationId:   migrationid,
		NextQueueName: nextQueueName,
	}

	s.GetSyncStore().Add(migrationid, sync)
}

func (s *syncer) FinishJob(migrationid string, role MigrationRole) error {
	obj, err := s.GetSyncStore().Fetch(migrationid)

	if err != nil {
		return err

	}

	syncObj, ok := obj.(Sync)

	if !ok {
		return errors.New("sync store did not get a sync object")
	}

	if role == MigrationRoleClient {
		syncObj.ClientDone = true
	} else {
		syncObj.ServerDone = true
	}

	// Notify RPC
	err = s.d.GetRPC().SendSyncNotification(migrationid, syncObj.CurrentStatus, role.PeersRole())

	if err != nil {
		return err
	}

	// Add it to the queue
	if syncObj.AllCompleted() {
		// Update the status of migration job
		obj, err = s.d.GetJobStore().Fetch(migrationid)

		if err != nil {
			return err
		}

		jobObj, ok := obj.(MigrationJob)

		if !ok {
			return errors.New("job store does not contain a migration job")
		}

		jobObj.Status = syncObj.NextStatus

		s.d.GetJobStore().Add(migrationid, jobObj)

		// Add the job back to the next specified queue if wanted
		if syncObj.NextQueueName != NullQueue {
			s.d.GetQueue(syncObj.NextQueueName).Push(migrationid)
		}
	}

	return nil
}

func (s *syncer) GetSyncStore() structures.Store {
	return s.syncstore
}

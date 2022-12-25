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
	WaitingStatus MigrationStatus

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
	// Marks the job with migrationid as complete. This will add the job to the next
	SyncStatus(migrationid string, callersRole MigrationRole, currentStatus MigrationStatus, nextStatus MigrationStatus, nextQueueName string) error

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

func NewSync(migrationid string, currentStatus MigrationStatus, nextStatus MigrationStatus) *Sync {
	return &Sync{
		ClientDone:    false,
		ServerDone:    false,
		WaitingStatus: nextStatus,
		MigrationId:   migrationid,
	}
}

// SyncStatus used for syncing stages of migration between server and clients. Both server and client must be
// done for migration to go to the new stage.
// This requires the migrationid, callers role, current and next status.
// If the syncing is successful meaning both server and client finnished processing, the job is added to the next queue.
// d
func (s *syncer) SyncStatus(migrationid string, callersRole MigrationRole, currentStatus MigrationStatus, nextStatus MigrationStatus, nextQueueName string) error {
	obj, err := s.GetSyncStore().Fetch(migrationid)

	if err != nil {
		return err
	}

	if syncObj, ok := obj.(Sync); !ok {
		return errors.New("sync-store did not get a sync object")
	} else {
		// Check if current status, if not same discard it with error.
		if syncObj.WaitingStatus != currentStatus {
			return errors.New("unknown current status")
		}

		// This means we have called this twice, no need to try to sync again.
		if syncObj.RoleCompleted(callersRole) {
			return nil
		}

		// Set the role completed and also notify the other peer.
		syncObj.SetCompleted(callersRole)

		// Notify the peer with RPC
		if !syncObj.AllCompleted() {
			// s.d.GetRPC().SyncNotification()
		} else {
			sync := NewSync(migrationid, currentStatus, nextStatus)

			// Update the sync store with new values
			s.GetSyncStore().Add(migrationid, sync)

			// Push the job to the next queue.
			s.d.GetQueue(nextQueueName).Push(migrationid)
		}

		return nil
	}

}

func (s *syncer) GetSyncStore() structures.Store {
	return s.syncstore
}

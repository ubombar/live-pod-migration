package daemon

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ubombar/live-pod-migration/pkg/migratord/structures"
)

// Denotes a sync object. This is used for waiting the other peer until both
// reach sync points.
type Sync struct {
	MigrationId string

	// // For client
	// ClientCurrent                MigrationStatus
	// ClientNext                   MigrationStatus
	// ClientStateCheckpointReached bool
	// ClientError                  error
	// ClientSuccessfull            bool

	// // For server
	// ServerCurrent                MigrationStatus
	// ServerNext                   MigrationStatus
	// ServerStateCheckpointReached bool
	// ServerError                  error
	// ServerSuccessfull            bool

	ClientStateFinished bool
	ServerStateFinished bool
	NextState           MigrationStatus
	Error               error
}

func (s *Sync) SetRoleStateFininshed(role MigrationRole) {
	if role == MigrationRoleClient {
		s.ClientStateFinished = true
	} else {
		s.ServerStateFinished = true
	}
}

func (s *Sync) StateCheckpointReached() bool {
	return s.ClientStateFinished && s.ServerStateFinished
}

func (s *Sync) StateCheckpointReachedSuccessfully() bool {
	return s.Error == nil && s.StateCheckpointReached()
}

func (s *Sync) GetNextState() MigrationStatus {
	if s.Error == nil {
		return s.NextState
	}

	return StatusError
}

// // By the role, marks the done segment true
// func (s *Sync) SetCompleted(role MigrationRole) {
// 	if role == MigrationRoleServer {
// 		s.ServerDone = true
// 	} else {
// 		s.ClientDone = true
// 	}
// }

// // If both lights are green, then it means current stage has finnished.
// func (s *Sync) AllCompletedSuccessfull() bool {
// 	return s.ClientDone && s.ServerDone
// }

// Return if the current role is completed
// func (s *Sync) RoleCompleted(role MigrationRole) bool {
// 	if role == MigrationRoleServer {
// 		return s.ServerDone
// 	} else {
// 		return s.ClientDone
// 	}
// }

type Syncer interface {
	// // Register the migration if with requested attributes. current status and the next job queue.
	RegisterJob(migrationid string) error

	// // Indicates the job has finnished. If the peer still processing it will not add the migration
	// // to the next queue until peer finishes.
	// FinishJob(migrationid string, role MigrationRole) error

	// // Get the stores
	GetSyncStore() structures.Store

	// // Notify in case of an error
	// FinishJobWithError(migrationid string, jobError error, role MigrationRole) error

	// New Versions
	// FinishStateAndSync(migrationid string, stateFinnished MigrationStatus, stateNext MigrationStatus, err error) error

	Prepare(migrationid string) error

	ConcludeState(migrationid string, nextState MigrationStatus) error
	ConcludeStateWithError(migrationid string, stateError error) error
}

type syncer struct {
	Syncer
	m         sync.Mutex
	d         Daemon
	syncstore structures.Store
}

func NewSyncer(daemon Daemon) *syncer {
	s := &syncer{
		m:         sync.Mutex{},
		d:         daemon,
		syncstore: structures.NewStore(MigrationSyncStore),
	}

	return s
}

func (s *syncer) GetSyncStore() structures.Store {
	return s.syncstore
}

func (s *syncer) RegisterJob(migrationid string) error {
	sync := Sync{
		MigrationId:         migrationid,
		ClientStateFinished: false,
		ServerStateFinished: false,
		NextState:           StatusPreparing,
		Error:               nil,
	}

	s.GetSyncStore().Add(migrationid, sync)

	return nil
}

func (s *syncer) ConcludeState(migrationid string, nextState MigrationStatus) error {
	obj, err := s.d.GetRoleStore().Fetch(migrationid)

	if err != nil {
		return err
	}

	role, ok := obj.(MigrationRole)

	if !ok {
		return errors.New("role store does not have migration role")
	}

	s.GetSyncStore().AtomicUpdate(migrationid, func(old interface{}, exists bool) (interface{}, bool) {
		if !exists {
			return nil, false
		}

		sync, ok := old.(Sync)

		if !ok {
			return nil, false
		}

		sync.SetRoleStateFininshed(role)
		sync.NextState = nextState
		sync.Error = nil

		return sync, true
	})

	s.d.GetQueue(SyncQueue).Push(migrationid)

	obj, _ = s.GetSyncStore().Fetch(migrationid)
	fmt.Printf("conclude: %v\n", obj)

	return s.d.GetRPC().NotifyPeerAboutStateChange(migrationid, nextState, nil)
}

func (s *syncer) ConcludeStateWithError(migrationid string, stateError error) error {
	obj, err := s.d.GetRoleStore().Fetch(migrationid)

	if err != nil {
		return err
	}

	role, ok := obj.(MigrationRole)

	if !ok {
		return errors.New("role store does not have migration role")
	}

	s.GetSyncStore().AtomicUpdate(migrationid, func(old interface{}, exists bool) (interface{}, bool) {
		if !exists {
			return nil, false
		}

		sync, ok := old.(Sync)

		if !ok {
			return nil, false
		}

		sync.SetRoleStateFininshed(role)
		sync.Error = stateError
		sync.NextState = StatusError

		return sync, true
	})

	// obj, _ := s.GetSyncStore().Fetch(migrationid)
	// fmt.Printf("obj: %v\n", obj)

	s.d.GetQueue(ErrorQueue).Push(migrationid)

	return s.d.GetRPC().NotifyPeerAboutStateChange(migrationid, StatusError, stateError)
}

func (s *syncer) Prepare(migrationid string) error {
	s.GetSyncStore().AtomicUpdate(migrationid, func(old interface{}, exists bool) (interface{}, bool) {
		if !exists {
			return nil, false
		}

		sync, ok := old.(Sync)

		if !ok {
			return nil, false
		}

		sync.ClientStateFinished = false
		sync.ServerStateFinished = false
		sync.NextState = StatusError
		sync.Error = nil

		return sync, true
	})

	obj, _ := s.GetSyncStore().Fetch(migrationid)
	fmt.Printf("prepare: %v\n", obj)

	return nil
}

// func (s *syncer) FinishStateAndSync(migrationid string, stateFinnished MigrationStatus, stateNext MigrationStatus, migrationError error) error {
// 	job, role, sync, err := s.getMigrationObjects(migrationid)

// 	return nil
// }

// func (s *syncer) RegisterJob(migrationid string, currentStatus MigrationStatus, nextQueueName string) error {
// 	sync := Sync{
// 		ClientDone:    false,
// 		ServerDone:    false,
// 		CurrentStatus: currentStatus,
// 		MigrationId:   migrationid,
// 		NextQueueName: nextQueueName,
// 	}

// 	s.GetSyncStore().Add(migrationid, sync)

// 	obj, err := s.d.GetJobStore().Fetch(migrationid)

// 	if err != nil {
// 		return err
// 	}

// 	jobObj, ok := obj.(MigrationJob)

// 	if !ok {
// 		return errors.New("job store does not contain a migration job")
// 	}

// 	jobObj.Status = currentStatus

// 	s.d.GetJobStore().Add(migrationid, jobObj)

// 	return nil
// }

// func (s *syncer) ChangeNextQueueName(migrationid string, nextQueueName string) error {
// 	obj, err := s.GetSyncStore().Fetch(migrationid)

// 	if err != nil {
// 		return err
// 	}

// 	syncObj, ok := obj.(Sync)

// 	if !ok {
// 		return errors.New("sync store did not get a sync object")
// 	}

// 	syncObj.NextQueueName = nextQueueName

// 	s.GetSyncStore().Add(migrationid, syncObj)

// 	return nil
// }

// func (s *syncer) FinishJob(migrationid string, role MigrationRole) error {
// 	s.m.Lock()
// 	defer s.m.Unlock()
// 	obj, err := s.GetSyncStore().Fetch(migrationid)

// 	if err != nil {
// 		return err
// 	}

// 	syncObj, ok := obj.(Sync)

// 	if !ok {
// 		return errors.New("sync store did not get a sync object")
// 	}

// 	obj, err = s.d.GetJobStore().Fetch(migrationid)

// 	if err != nil {
// 		return err
// 	}

// 	jobObj, ok := obj.(MigrationJob)

// 	if !ok {
// 		return errors.New("sync store did not get a sync object")
// 	}

// 	// Return if role is already completed
// 	if syncObj.RoleCompleted(role) {
// 		return nil
// 	} else {
// 		syncObj.SetCompleted(role)
// 		s.GetSyncStore().Add(migrationid, syncObj)

// 		// Notify RPC
// 		err = s.d.GetRPC().SendSyncNotification(migrationid, syncObj.CurrentStatus, role.PeersRole(), jobObj.Error)

// 		if err != nil {
// 			return err
// 		}
// 	}

// 	// Add it to the queue
// 	if syncObj.AllCompleted() {
// 		// Add the job back to the next specified queue if wanted
// 		if syncObj.NextQueueName != NullQueue {
// 			s.d.GetQueue(syncObj.NextQueueName).Push(migrationid)
// 		} else {
// 			obj, err := s.d.GetJobStore().Fetch(migrationid)

// 			if err != nil {
// 				return err
// 			}

// 			jobObj, ok := obj.(MigrationJob)

// 			if !ok {
// 				return errors.New("job store does not contain a migration job")
// 			}

// 			if jobObj.Error == nil {
// 				jobObj.Status = StatusDone
// 			}

// 			s.d.GetJobStore().Add(migrationid, jobObj)
// 		}
// 		s.GetSyncStore().Delete(migrationid)
// 	}

// 	return nil
// }

// // Calls the FinishJob but changes the nextQueue to error queue and job.Error to the given error.
// func (s *syncer) FinishJobWithError(migrationid string, jobError error, role MigrationRole) error {
// 	s.m.Lock()
// 	obj, err := s.GetSyncStore().Fetch(migrationid)

// 	if err != nil {
// 		return err
// 	}

// 	syncObj, ok := obj.(Sync)

// 	if !ok {
// 		return errors.New("sync store did not get a sync object")
// 	}

// 	// Set the next queue to error queue, if it is not null queue.
// 	if syncObj.NextQueueName != NullQueue {
// 		syncObj.NextQueueName = ErrorQueue
// 	}

// 	// Update the syncObject
// 	s.GetSyncStore().Add(migrationid, syncObj)

// 	obj, err = s.d.GetJobStore().Fetch(migrationid)

// 	if err != nil {
// 		return err
// 	}

// 	jobObj, ok := obj.(MigrationJob)

// 	if !ok {
// 		return errors.New("sync store did not get a sync object")
// 	}

// 	jobObj.Error = jobError
// 	jobObj.Status = StatusError

// 	// Update the job object
// 	s.d.GetJobStore().Add(migrationid, jobObj)

// 	s.m.Unlock()
// 	return s.FinishJob(migrationid, role)
// }

// func (s *syncer) GetSyncStore() structures.Store {
// 	return s.syncstore
// }

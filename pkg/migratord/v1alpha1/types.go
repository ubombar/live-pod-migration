package v1alpha1

import (
	"errors"
	"sync"
	"time"
)

const (
	DefaultMigrationQueueCapacity = 16
)

const DefaultCheckpointDirectory = "/home/ubombar/migratord"

type MigrationStatus string

const (
	Preparing     MigrationStatus = "Preparing"
	Checkpointing MigrationStatus = "Checkpointing"
	Transfering   MigrationStatus = "Transfering"
	Restoring     MigrationStatus = "Restoring"
	Done          MigrationStatus = "Done"
	Error         MigrationStatus = "Error"
)

type MigrationMethod string

const (
	Basic    MigrationMethod = "Basic"
	Precopy  MigrationMethod = "Precopy"
	Postcopy MigrationMethod = "Postcopy"
)

type MigrationRole string

const (
	RoleClient MigrationRole = "Client"
	RoleServer MigrationRole = "Server"
)

type MigrationInfo interface {
	Type() MigrationMethod
}

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

	// Contains additional information about the migration
	MigrationInfo MigrationInfo
}

type MigrationQueue struct {
	mutex sync.Mutex
	queue chan *MigrationJob
}

func NewMigrationQueue(maxLength int) (*MigrationQueue, error) {
	if maxLength <= 0 {
		return nil, errors.New("queue length cannot be zero or negative")
	}

	queue := &MigrationQueue{
		mutex: sync.Mutex{},
		queue: make(chan *MigrationJob, maxLength),
	}

	return queue, nil
}

func (q *MigrationQueue) Push(m *MigrationJob) bool {
	if len(q.queue) == cap(q.queue) {
		return false
	}

	q.queue <- m
	return true
}

// Blocking
func (q *MigrationQueue) Pop() *MigrationJob {
	m := <-q.queue
	return m
}

func (q *MigrationQueue) PopNonBlock() (*MigrationJob, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.queue) == 0 {
		return nil, false
	}

	m := <-q.queue
	return m, true
}

type MigrationMap struct {
	mutex sync.Mutex
	mmap  map[string]*MigrationJob
}

func NewMigrationMap() (*MigrationMap, error) {
	mmap := &MigrationMap{
		mutex: sync.Mutex{},
		mmap:  make(map[string]*MigrationJob),
	}

	return mmap, nil
}

func (q *MigrationMap) Get(migrationId string) (*MigrationJob, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	migration, ok := q.mmap[migrationId]
	return migration, ok
}

func (q *MigrationMap) Save(m *MigrationJob) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.mmap[m.MigrationId] = m
}

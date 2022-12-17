package v1alpha1

import (
	"errors"
	"sync"
	"time"
)

const (
	DefaultMigrationQueueCapacity = 16
)

const (
	Pending       = "Peding"
	Checkpointing = "Checkpointing"
	Transfering   = "Transfering"
	Restoring     = "Restoring"
	Done          = "Done"
	Error         = "Error"
)

type MigrationStatus string

const (
	RoleClient = "Role Client"
	RoleServer = "Role Server"
)

type MigrationRole string

type MigrationJob struct {
	// Used for describing the migration
	MigrationId string

	// Client and Servers ip address
	ClientIP string
	ServerIP string

	// Cotnainer identification number
	ContainerID string

	// Migration status describes in which stage migration is
	Status MigrationStatus

	// Shows if the container is running
	Running bool

	// Creation date
	CreationDate time.Time

	// Role of the migratord that owns this object
	Role MigrationRole
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

func (q *MigrationMap) Save(m *MigrationJob) bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if _, ok := q.mmap[m.MigrationId]; ok {
		return false
	}

	q.mmap[m.MigrationId] = m

	return true
}

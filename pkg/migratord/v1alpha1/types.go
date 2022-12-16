package v1alpha1

import (
	"errors"
	"sync"
	"time"
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

type Migration struct {
	// Used for describing the migration
	MigrationId string

	// Client and Servers ip address
	ClientIP string
	ServerIP string

	// Cotnainer identification number
	ContainerID string

	// Migration status describes in which stage migration is
	Status MigrationStatus

	// Shows if the container is running.
	Running bool

	// Creation date
	CreationDate time.Time
}

type MigrationQueue struct {
	mutex sync.Mutex
	queue chan Migration
}

func NewMigrationQueue(maxLength int) (*MigrationQueue, error) {
	if maxLength <= 0 {
		return nil, errors.New("queue length cannot be zero or negative")
	}

	queue := &MigrationQueue{
		mutex: sync.Mutex{},
		queue: make(chan Migration, maxLength),
	}

	return queue, nil
}

func (q *MigrationQueue) Push(m Migration) bool {

}

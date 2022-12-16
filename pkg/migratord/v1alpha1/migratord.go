package v1alpha1

import (
	"context"
	"errors"

	"github.com/docker/docker/client"
)

// This will run as a system deamon. Checks incoming client and server migration requests.
type Migratord struct {
	Client           *client.Client
	OSType           string
	ClientAPIVersion string

	// Incoming queue
	IncomingMigrations *MigrationQueue
}

func NewMigratord() (*Migratord, error) {
	cl, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return nil, err
	}

	ping, err := cl.Ping(context.Background())

	if err != nil {
		return nil, err
	}

	if !ping.Experimental {
		return nil, errors.New("experimental mode is not eabled on node")
	}

	queue, err := NewMigrationQueue(DefaultMigrationQueueCapacity)

	if err != nil {
		return nil, err
	}

	migratord := &Migratord{
		Client:             cl,
		OSType:             ping.OSType,
		ClientAPIVersion:   ping.APIVersion,
		IncomingMigrations: queue,
	}

	return migratord, nil
}

// Listens for incoming requests. If the request comes from another migratord then it will act like a server.
// If it comes from someone else, then it will act tike a client.
func (m *Migration) Run() {
	// Wait for a gRPC connection,
}

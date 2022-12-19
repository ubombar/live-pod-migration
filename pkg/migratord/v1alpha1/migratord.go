package v1alpha1

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/docker/docker/client"
	pb "github.com/ubombar/live-pod-migration/pkg/migrator"
	"google.golang.org/grpc"
)

// This will run as a system deamon. Checks incoming client and server migration requests.
type Migratord struct {
	Client           *client.Client
	OSType           string
	ClientAPIVersion string
	Address          string
	Port             int

	// These are handlers for client and server roles in migration.
	serverHandler *serverMigationHandler
	clientHandler *clientMigationHandler

	// Incoming queue, for managing migrations actively. Only used by client role migrator.
	MigrationQueue *MigrationQueue

	// This is for passively and actively managed migration.
	MigrationMap *MigrationMap
}

func NewMigratord(address string, port int) (*Migratord, error) {
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

	mmap, err := NewMigrationMap()

	if err != nil {
		return nil, err
	}

	serverHandler := &serverMigationHandler{}
	clientHandler := &clientMigationHandler{}

	migratord := &Migratord{
		Client:           cl,
		OSType:           ping.OSType,
		ClientAPIVersion: ping.APIVersion,
		MigrationQueue:   queue,
		MigrationMap:     mmap,
		Address:          address,
		Port:             port,
		serverHandler:    serverHandler,
		clientHandler:    clientHandler,
	}

	// Set the migratord pointers, so it will be accessible inside handlers.
	serverHandler.parent = migratord
	clientHandler.parent = migratord

	return migratord, nil
}

// Listens for incoming requests. If the request comes from another migratord then it will act like a server.
// If it comes from someone else, then it will act tike a client.
func (m *Migratord) Run() {
	// Run the consumer queue, if gets a new migration then spawn a new go routing.
	go func() {
		// Always consume from the queue
		for {
			migration := m.MigrationQueue.Pop()
			m.MigrationMap.Save(migration)
			go m.clientHandler.PerformMigration(migration)
		}
	}()
	// Before listening this will
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", m.Address, m.Port))

	if err != nil {
		return
	}

	server := grpc.NewServer(grpc.EmptyServerOption{})
	pb.RegisterMigratorServiceServer(server, m.serverHandler)
	server.Serve(lis)
}

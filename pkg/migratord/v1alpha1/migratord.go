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

	handler *MigratordRPC

	// Incoming queue
	IncomingMigrations *MigrationQueue
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

	handler := &MigratordRPC{}

	migratord := &Migratord{
		Client:             cl,
		OSType:             ping.OSType,
		ClientAPIVersion:   ping.APIVersion,
		IncomingMigrations: queue,
		Address:            address,
		Port:               port,
		handler:            handler,
	}

	// Set the inverse pointer so migratord will be accessible from hane gRPC handler.
	handler.parent = migratord

	return migratord, nil
}

type MigratordRPC struct {
	pb.MigratorServiceServer
	parent *Migratord
}

// Listens for incoming requests. If the request comes from another migratord then it will act like a server.
// If it comes from someone else, then it will act tike a client.
func (m *Migratord) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", m.Address, m.Port))

	if err != nil {
		return
	}

	server := grpc.NewServer(grpc.EmptyServerOption{})
	pb.RegisterMigratorServiceServer(server, m.handler)
	server.Serve(lis)
}

func (m *MigratordRPC) CreateMigrationJob(ctx context.Context, req *pb.CreateMigrationJobRequest) (*pb.CreateMigrationJobResponse, error) {
	if req == nil {
		return nil, errors.New("incoming request is nil")
	}

	// Preflight checks
	containerJSON, err := m.parent.Client.ContainerInspect(ctx, req.ContainerId)

	if err != nil {
		return nil, err
	}

	if !containerJSON.State.Running {
		return nil, errors.New("container is not running")
	}

	// Assuming everything is perfect and we are able to share this migration information with out peer migrator.

	return nil, nil
}

func (m *MigratordRPC) ShareMigrationJob(ctx context.Context, req *pb.ShareMigrationJobRequest) (*pb.ShareMigrationJobResponse, error) {
	if req == nil {
		return nil, errors.New("incoming request is nil")
	}

	return nil, nil
}

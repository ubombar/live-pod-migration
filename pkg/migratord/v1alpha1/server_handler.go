package v1alpha1

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	pb "github.com/ubombar/live-pod-migration/pkg/migrator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serverMigationHandler struct {
	pb.MigratorServiceServer
	parent *Migratord
}

// Tells any migratord to get in client role.
func (m *serverMigationHandler) CreateMigrationJob(ctx context.Context, req *pb.CreateMigrationJobRequest) (*pb.CreateMigrationJobResponse, error) {
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
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", req.PeerAddress, req.PeerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, errors.New("cannot reach peer")
	}

	client := pb.NewMigratorServiceClient(conn)

	resp, err := client.ShareMigrationJob(ctx, &pb.ShareMigrationJobRequest{
		PeerAddress:    m.parent.Address,
		PeerPort:       int32(m.parent.Port),
		ContainerId:    containerJSON.Image,
		ContainerImage: containerJSON.Image,
		ContainerName:  containerJSON.Name,
	})

	if err != nil || !resp.Accepted {
		return nil, errors.New("peer does not accept the migration job")
	}

	// Create the migration object
	migObject := &Migration{
		ClientIP:     req.PeerAddress,
		ServerIP:     m.parent.Address,
		MigrationId:  resp.MigrationId,
		ContainerID:  req.ContainerId,
		Status:       Pending,
		Running:      true,
		CreationDate: time.Unix(resp.CreatonUnixTime, 0),
		Role:         RoleClient,
	}

	// Add the migration to the queue
	m.parent.MigrationQueue.Push(migObject)

	return &pb.CreateMigrationJobResponse{Accepted: true}, nil
}

// Migratord with client role invokes it's peer. If works it's peer gets in a server role.
func (m *serverMigationHandler) ShareMigrationJob(ctx context.Context, req *pb.ShareMigrationJobRequest) (*pb.ShareMigrationJobResponse, error) {
	if req == nil {
		return nil, errors.New("incoming request is nil")
	}

	images, err := m.parent.Client.ImageList(ctx, types.ImageListOptions{})

	if err != nil {
		return nil, err
	}

	containsImage := false

	for _, image := range images {
		if image.ID == req.ContainerImage {
			containsImage = true
			break
		}
	}

	if !containsImage {
		fmt.Println("Warning, the server migrator doesn't have the specified image!")
	}

	// Create the migrationid from migration string, might need work here...
	creationDate := time.Now()
	createDateString := fmt.Sprint(creationDate.Nanosecond())
	hash := sha256.New()
	hash.Write([]byte(req.ContainerId))
	hash.Write([]byte(req.ContainerImage))
	byteArray := hash.Sum([]byte(createDateString))
	migrationId := fmt.Sprintf("%x", byteArray)

	// Create the migration object
	migObject := &Migration{
		ClientIP:     req.PeerAddress,
		ServerIP:     m.parent.Address,
		MigrationId:  migrationId,
		ContainerID:  req.ContainerId,
		Status:       Pending,
		Running:      true,
		CreationDate: creationDate,
		Role:         RoleServer,
	}

	// Add the migration to the migration map
	m.parent.MigrationMap.Save(migObject)

	return &pb.ShareMigrationJobResponse{
		Accepted:        true,
		MigrationId:     migrationId,
		CreatonUnixTime: creationDate.Unix(),
	}, nil
}

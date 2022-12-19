package v1alpha1

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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

	logrus.Println("New migration job received.")

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
		ContainerId:    containerJSON.ID,
		ContainerImage: containerJSON.Image,
		ContainerName:  containerJSON.Name,
	})

	if err != nil || !resp.Accepted {
		return nil, errors.New("peer does not accept the migration job")
	}

	var migrationMethod MigrationMethod

	switch req.Method {
	case pb.MigrationMethod_Basic:
		migrationMethod = Basic
	case pb.MigrationMethod_Precopy:
		migrationMethod = Precopy
	case pb.MigrationMethod_Postcopy:
		migrationMethod = Postcopy
	}

	// Create the migration object
	migObject := &MigrationJob{
		ClientIP:     req.PeerAddress,
		ServerIP:     m.parent.Address,
		MigrationId:  resp.MigrationId,
		ContainerID:  containerJSON.ID,
		Status:       Preparing,
		Running:      true,
		CreationDate: time.Unix(resp.CreatonUnixTime, 0),
		Role:         RoleClient,
		Method:       migrationMethod,
		ServerPort:   int(req.PeerPort),
		ClientPort:   m.parent.Port,
	}

	// Add the migration to the queue
	m.parent.MigrationQueue.Push(migObject)

	logrus.Printf("New migration accepted as client. Migration id %s\n", resp.MigrationId)

	response := &pb.CreateMigrationJobResponse{
		Accepted:        true,
		MigrationId:     resp.MigrationId,
		CreatonUnixTime: migObject.CreationDate.Unix(),
	}

	return response, nil
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

	containerInspect, err := m.parent.Client.ContainerInspect(ctx, req.ContainerId)

	if err != nil {
		return nil, err
	}

	// Create the migrationid from migration string.
	creationDate := time.Now()
	migrationId := strings.ReplaceAll(uuid.New().String(), "-", "")

	var migrationMethod MigrationMethod

	switch req.Method {
	case pb.MigrationMethod_Basic:
		migrationMethod = Basic
	case pb.MigrationMethod_Precopy:
		migrationMethod = Precopy
	case pb.MigrationMethod_Postcopy:
		migrationMethod = Postcopy
	}

	// Create the migration object
	migObject := &MigrationJob{
		ClientIP:     req.PeerAddress,
		ServerIP:     m.parent.Address,
		MigrationId:  migrationId,
		ContainerID:  containerInspect.ID,
		Status:       Preparing,
		Running:      true,
		CreationDate: creationDate,
		Role:         RoleServer,
		Method:       migrationMethod,
		ClientPort:   int(req.PeerPort),
		ServerPort:   m.parent.Port,
	}

	// Add the migration to the migration map
	m.parent.MigrationMap.Save(migObject)

	logrus.Printf("New migration accepted as server. Migration id %s\n", migrationId)

	return &pb.ShareMigrationJobResponse{
		Accepted:        true,
		MigrationId:     migrationId,
		CreatonUnixTime: creationDate.Unix(),
	}, nil
}

// Updates the status of the migration, invoked in server.
func (m *serverMigationHandler) UpdateMigrationStatus(ctx context.Context, req *pb.UpdateMigrationStatusRequest) (*pb.UpdateMigrationStatusResponse, error) {
	if req == nil {
		return nil, errors.New("incoming request is nil")
	}

	migrationJob, ok := m.parent.MigrationMap.Get(req.MigrationId)

	if !ok {
		logrus.Errorln("cannot find given migration")
		return nil, errors.New("cannot find migration with given id")
	}

	// Maybe add a checker for this before changing the status
	migrationJob.Status = MigrationStatus(req.NewStatus)
	migrationJob.Running = req.NewRunning

	m.parent.MigrationMap.Save(migrationJob)

	return &pb.UpdateMigrationStatusResponse{}, nil
}

// Gets the status of the migration, invoked in server.
func (m *serverMigationHandler) GetMigrationStatus(ctx context.Context, req *pb.GetMigrationStatusRequest) (*pb.GetMigrationStatusResponse, error) {
	if req == nil {
		return nil, errors.New("incoming request is nil")
	}

	job, ok := m.parent.MigrationMap.Get(req.MigrationId)

	if !ok {
		return nil, errors.New("cannot find migration with given id")
	}

	resp := &pb.GetMigrationStatusResponse{
		MigrationId:     job.MigrationId,
		ServerIp:        job.ServerIP,
		ClientIp:        job.ClientIP,
		ContainerId:     job.ContainerID,
		MigrationStatus: string(job.Status),
		Running:         job.Running,
		MigrationTime:   job.CreationDate.Unix(),
		MigrationRole:   string(job.Role),
		MigrationMethod: string(job.Method),
	}

	return resp, nil
}

// Stream the checkpoint file
func (m *serverMigationHandler) SendViaSCP(ctx context.Context, req *pb.SendViaSCPRequest) (*pb.SendViaSCPResponse, error) {
	if req == nil {
		return nil, errors.New("incoming request is nil")
	}

	return nil, nil
}

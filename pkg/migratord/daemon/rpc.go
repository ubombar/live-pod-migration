package daemon

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	pb "github.com/ubombar/live-pod-migration/pkg/generated"
	"google.golang.org/grpc"
)

type RPCConfig struct {
	Address string
	Port    int
}

type RPCPeer struct {
	Address string
	Port    int
}

func (r RPCPeer) String() string {
	return fmt.Sprintf("%v:%d", r.Address, r.Port)
}

type RPC interface {
	Run() error
	// PeerShareMigrationJob(RPCPeer, *pb.CreateMigrationJobRequest) (*pb.CreateMigrationJobResponse, error)

}

type rpc struct {
	RPC
	pb.UnimplementedMigratorServiceServer
	config RPCConfig
	d      Daemon
}

func NewRPC(config RPCConfig, daemon Daemon) *rpc {
	r := &rpc{
		config: config,
		d:      daemon,
	}

	return r
}

func (r *rpc) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", r.config.Address, r.config.Port))

	if err != nil {
		return err
	}

	server := grpc.NewServer(grpc.EmptyServerOption{})
	pb.RegisterMigratorServiceServer(server, r)

	go func() {
		server.Serve(lis)
	}()

	return nil
}

func (r *rpc) CreateMigrationJob(ctx context.Context, req *pb.CreateMigrationJobRequest) (*pb.CreateMigrationJobResponse, error) {
	peerReq := &pb.ShareMigrationJobRequest{
		ContainerId:            req.ContainerId,
		ClientContainerRuntime: req.ClientContainerRuntime,
		ServerContainerRuntime: req.ServerContainerRuntime,
		ClientAddress:          r.d.GetConfig().SelfAddress,
		ClientPort:             int32(r.d.GetConfig().SelfPort),
		ServerKey:              req.ServerKey,
		ServerUser:             req.ServerUser,
		Method:                 req.Method,
	}

	peerResp, err := r.ShareMigrationJob(ctx, peerReq)

	if err != nil {
		logrus.Warnln("cannot process migration job")
		return nil, err
	}

	// Create a new job and add it to queue
	job := &MigrationJob{
		MigrationId:       peerResp.MigrationId,
		ClientIP:          r.d.GetConfig().SelfAddress,
		ClientPort:        r.d.GetConfig().SelfPort,
		ServerIP:          req.ServerAddress,
		ServerPort:        int(req.ServerPort),
		ClientContainerID: req.ContainerId,
		ServerContainerID: "",
		Status:            StatusPreparing,
		Running:           true,
		PrivateKey:        req.ServerKey,
		Username:          req.ServerUser,
		CreationDate:      time.Unix(peerResp.CreatonUnixTime, 0),
		Role:              MigrationRoleClient,
		Method:            MigrationMethod(req.Method),
	}

	// Register the job to store
	r.d.GetJobStore().Add(job.MigrationId, &job)
	r.d.GetRoleStore().Add(job.MigrationId, job.Role)

	// Add the job to the queue
	r.d.GetQueue(IncomingQueue).Push(job.MigrationId)

	return &pb.CreateMigrationJobResponse{
		MigrationId:     peerResp.MigrationId,
		CreatonUnixTime: peerResp.CreatonUnixTime,
	}, nil
}

func (r *rpc) ShareMigrationJob(ctx context.Context, req *pb.ShareMigrationJobRequest) (*pb.ShareMigrationJobResponse, error) {
	uuidHash, _ := uuid.NewRandom()
	migrationId := fmt.Sprint(uuidHash.String())

	// Create a new job and add it to queue
	job := &MigrationJob{
		MigrationId:       migrationId,
		ServerIP:          r.d.GetConfig().SelfAddress,
		ServerPort:        r.d.GetConfig().SelfPort,
		ClientIP:          req.ClientAddress,
		ClientPort:        int(req.ClientPort),
		ClientContainerID: req.ContainerId,
		ServerContainerID: "",
		Status:            StatusPreparing,
		Running:           true,
		PrivateKey:        req.ServerKey,
		Username:          req.ServerUser,
		CreationDate:      time.Now(),
		Role:              MigrationRoleClient,
		Method:            MigrationMethod(req.Method),
	}

	// Register the job to store
	r.d.GetJobStore().Add(job.MigrationId, &job)
	r.d.GetRoleStore().Add(job.MigrationId, job.Role)

	// Add the job to the queue
	r.d.GetQueue(IncomingQueue).Push(job.MigrationId)

	return &pb.ShareMigrationJobResponse{
		MigrationId:     migrationId,
		CreatonUnixTime: job.CreationDate.Unix(),
	}, nil
}

// Invoked on sync notification
func (r *rpc) SyncNotification(ctx context.Context, req *pb.SyncNotificationRequest) (*pb.SyncNotificationResponse, error) {
	obj, err := r.d.GetSyncer().GetSyncStore().Fetch(req.MigrationId)

	if err != nil {
		return nil, err
	}

	syncObj, ok := obj.(Sync)

	if !ok {
		return nil, errors.New("sync store does not contain sync object")
	}

	role, err := r.d.GetRoleStore().Fetch(req.MigrationId)

	if err != nil {
		return nil, err
	}

	var peersRole MigrationRole

	// Find our peers role
	if role == MigrationRoleClient {
		peersRole = MigrationRoleServer
	} else {
		peersRole = MigrationRoleClient
	}

	// If they are not waiting for the same state.
	if syncObj.WaitingStatus != MigrationStatus(req.CurrentStateName) {
		return nil, errors.New("client and server does not wait for the same state")
	}

	// Mark peers job as done
	syncObj.SetCompleted(peersRole)

	// If client and server has completed, then add it to the next queue.
	if syncObj.AllCompleted() {
		// Now create a new sync and add it to the qeueue
		sync := NewSync(req.MigrationId, MigrationStatus(req.CurrentStateName), MigrationStatus(req.NextStateName))

		// Update the sync store with new values
		r.d.GetSyncer().GetSyncStore().Add(req.MigrationId, sync)

		// Push the job to the next queue.
		r.d.GetQueue(req.NextQueueName).Push(req.MigrationId)
	}

	return &pb.SyncNotificationResponse{}, nil
}

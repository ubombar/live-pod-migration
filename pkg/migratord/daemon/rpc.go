package daemon

import (
	"context"
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
	daemon Daemon
}

func NewRPC(config RPCConfig, daemon Daemon) *rpc {
	r := &rpc{
		config: config,
		daemon: daemon,
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
		ClientAddress:          r.daemon.GetConfig().SelfAddress,
		ClientPort:             int32(r.daemon.GetConfig().SelfPort),
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
		ClientIP:          r.daemon.GetConfig().SelfAddress,
		ClientPort:        r.daemon.GetConfig().SelfPort,
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
	r.daemon.GetJobStore().Add(job.MigrationId, &job)
	r.daemon.GetRoleStore().Add(job.MigrationId, job.Role)

	// Add the job to the queue
	r.daemon.GetQueue(IncomingQueue).Push(job.MigrationId)

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
		ServerIP:          r.daemon.GetConfig().SelfAddress,
		ServerPort:        r.daemon.GetConfig().SelfPort,
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
	r.daemon.GetJobStore().Add(job.MigrationId, &job)
	r.daemon.GetRoleStore().Add(job.MigrationId, job.Role)

	// Add the job to the queue
	r.daemon.GetQueue(IncomingQueue).Push(job.MigrationId)

	return &pb.ShareMigrationJobResponse{
		MigrationId:     migrationId,
		CreatonUnixTime: job.CreationDate.Unix(),
	}, nil
}

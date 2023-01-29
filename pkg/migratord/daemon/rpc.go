package daemon

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	pb "github.com/ubombar/live-pod-migration/pkg/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	SendSyncNotification(migrationid string, currentStatus MigrationStatus, peersRole MigrationRole, jobError error) error
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

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", req.ServerAddress, req.ServerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	client := pb.NewMigratorServiceClient(conn)
	peerResp, err := client.ShareMigrationJob(ctx, peerReq)

	logrus.Infof("%v:%v\n", req.ServerAddress, req.ServerPort)

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
	r.d.GetJobStore().Add(job.MigrationId, *job)
	r.d.GetRoleStore().Add(job.MigrationId, MigrationRoleClient)

	// Add the job to the queue
	r.d.GetQueue(IncomingQueue).Push(job.MigrationId)

	return &pb.CreateMigrationJobResponse{
		MigrationId:     peerResp.MigrationId,
		CreatonUnixTime: peerResp.CreatonUnixTime,
	}, nil
}

func (r *rpc) ShareMigrationJob(ctx context.Context, req *pb.ShareMigrationJobRequest) (*pb.ShareMigrationJobResponse, error) {
	uuidHash, _ := uuid.NewRandom()
	migrationId := fmt.Sprint(strings.ReplaceAll(uuidHash.String(), "-", ""))

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
		Role:              MigrationRoleServer,
		Method:            MigrationMethod(req.Method),
	}

	// Register the job to store
	r.d.GetJobStore().Add(job.MigrationId, *job)
	r.d.GetRoleStore().Add(job.MigrationId, MigrationRoleServer)

	// Add the job to the queue
	r.d.GetQueue(IncomingQueue).Push(job.MigrationId)

	return &pb.ShareMigrationJobResponse{
		MigrationId:     migrationId,
		CreatonUnixTime: job.CreationDate.Unix(),
	}, nil
}

// Invoked on sync notification
func (r *rpc) SyncNotification(ctx context.Context, req *pb.SyncNotificationRequest) (*pb.SyncNotificationResponse, error) {
	// Get the migration role
	obj, err := r.d.GetRoleStore().Fetch(req.MigrationId)

	if err != nil {
		return nil, err
	}

	roleObj, ok := obj.(MigrationRole)

	if !ok {
		return nil, errors.New("role store does not contain role")
	}

	// Invoke finnish job function of syncer with peers role. If finished with error, invoke FinishJobWithError
	if req.FinishedSuccessful {
		r.d.GetSyncer().FinishJob(req.MigrationId, roleObj.PeersRole())
	} else {
		r.d.GetSyncer().FinishJobWithError(req.MigrationId, errors.New(req.ErrorMessage), roleObj.PeersRole())
	}

	return &pb.SyncNotificationResponse{}, nil
}

func (r *rpc) SendSyncNotification(migrationid string, statusFinished MigrationStatus, peersRole MigrationRole, jobError error) error {
	obj, err := r.d.GetJobStore().Fetch(migrationid)

	if err != nil {
		return err
	}

	jobObj, ok := obj.(MigrationJob)

	if !ok {
		return errors.New("job store does not contain a job object")
	}

	address := jobObj.AddressString(peersRole)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	client := pb.NewMigratorServiceClient(conn)

	errorMessage := ""
	finishedSuccessful := jobError == nil

	if !finishedSuccessful {
		errorMessage = jobError.Error()
	}

	_, err = client.SyncNotification(context.Background(), &pb.SyncNotificationRequest{
		MigrationId:        migrationid,
		StateFinished:      string(statusFinished),
		ErrorMessage:       errorMessage,
		FinishedSuccessful: finishedSuccessful,
	})

	if err != nil {
		return err
	}

	return nil
}

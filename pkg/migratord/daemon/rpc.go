package daemon

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sort"
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
	// SendSyncNotification(migrationid string, currentStatus MigrationStatus, peersRole MigrationRole, jobError error) error

	NotifyPeerAboutStateChange(migrationid string, finnished MigrationStatus, next MigrationStatus, jobError error) error
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

func (r *rpc) InformStateChange(ctx context.Context, req *pb.InformStateChangeRequest) (*pb.InformStateChangeResponse, error) {
	job, role, _, err := FromMigrationId(req.MigrationId, r.d)

	if err != nil {
		return nil, err
	}

	r.d.GetSyncer().GetSyncStore().AtomicUpdate(req.MigrationId, func(old interface{}, exists bool) (interface{}, bool) {
		if !exists {
			return nil, false
		}

		return nil, false
	})

	return nil, nil
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

	// rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing dial tcp 192.168.122.1:4545: connect: connection refused"

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

	// time.Sleep(1000 * time.Second)

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

	// logrus.Infoln("received! sync notification")
	logrus.Info("Received SyncNotification\n")

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

func (r *rpc) GetMigrationJob(ctx context.Context, req *pb.GetMigrationJobRequest) (*pb.GetMigrationJobResponse, error) {
	jobs := []*pb.MigrationJob{}

	keys := r.d.GetJobStore().Keys()

	sort.Strings(keys)

	for _, key := range keys {
		obj, err := r.d.GetJobStore().Fetch(key)

		if err != nil {
			continue
		}

		job, ok := obj.(MigrationJob)

		if !ok {
			continue
		}

		obj, err = r.d.GetRoleStore().Fetch(key)

		if err != nil {
			continue
		}

		role, ok := obj.(MigrationRole)

		if !ok {
			continue
		}

		errorString := ""

		if job.Error != nil {
			errorString = job.Error.Error()
		}

		jobs = append(jobs, &pb.MigrationJob{
			MigrationId:     job.MigrationId,
			ClientAddress:   job.ClientIP,
			ServerAddress:   job.ServerIP,
			ClientPort:      int32(job.ClientPort),
			ServerPort:      int32(job.ServerPort),
			CotninerId:      job.ClientContainerID,
			MigrationStatus: string(job.Status),
			ErrorString:     errorString,
			Running:         job.Running,
			CreationDate:    strings.Split(job.CreationDate.String(), "+")[0],
			Role:            string(role),
			MigrationMethod: string(job.Method),
		})
	}

	return &pb.GetMigrationJobResponse{Jobs: jobs}, nil
}

func (r *rpc) NotifyPeerAboutStateChange(migrationid string, finnished MigrationStatus, next MigrationStatus, jobError error) error {
	job, role, _, err := FromMigrationId(migrationid, r.d)
	if err != nil {
		return err
	}

	peerAddress := job.AddressString(role.PeersRole())
	conn, err := grpc.Dial(peerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return err
	}

	client := pb.NewMigratorServiceClient(conn)

	informStateChangeRequest := &pb.InformStateChangeRequest{
		MigrationId:   migrationid,
		FinishedState: string(finnished),
		NextState:     string(next),
		ErrorString:   nil,
	}

	if jobError != nil {
		errorString := jobError.Error()
		informStateChangeRequest.ErrorString = &errorString
	}

	_, err = client.InformStateChange(context.Background(), informStateChangeRequest)

	if err != nil {
		logrus.Error("error occured while informing peer for state change")
		return err
	}

	return nil
}

// func (r *rpc) SendSyncNotification(migrationid string, statusFinished MigrationStatus, peersRole MigrationRole, jobError error) error {
// 	obj, err := r.d.GetJobStore().Fetch(migrationid)

// 	if err != nil {
// 		return err
// 	}

// 	jobObj, ok := obj.(MigrationJob)

// 	if !ok {
// 		return errors.New("job store does not contain a job object")
// 	}

// 	address := jobObj.AddressString(peersRole)
// 	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

// 	if err != nil {
// 		return err
// 	}

// 	client := pb.NewMigratorServiceClient(conn)

// 	errorMessage := ""
// 	finishedSuccessful := jobError == nil

// 	if !finishedSuccessful {
// 		errorMessage = jobError.Error()
// 	}

// 	for i := 0; i < 5; i++ {
// 		// logrus.Info("Sending SyncNotification\n")
// 		_, err = client.SyncNotification(context.Background(), &pb.SyncNotificationRequest{
// 			MigrationId:        migrationid,
// 			StateFinished:      string(statusFinished),
// 			ErrorMessage:       errorMessage,
// 			FinishedSuccessful: finishedSuccessful,
// 		})
// 		// logrus.Info("Sent SyncNotification\n")

// 		if err != nil {
// 			logrus.Info("Error sending sync notification retrying\n")
// 			time.Sleep(1 * time.Second)
// 			continue
// 		}

// 		logrus.Info("Sent sync notification!")

// 		break

// 	}

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

package daemon

import (
	"context"
	"errors"
	"fmt"
	"net"

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
	PeerCreateMigrationJob(RPCPeer, pb.CreateMigrationJobRequest) (*pb.CreateMigrationJobResponse, error)
}

type rpc struct {
	RPC
	pb.UnimplementedMigratorServiceServer
	config RPCConfig
	daemon *Daemon
}

func NewRPC(config RPCConfig, daemon *Daemon) *rpc {
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
	if req.Source == pb.Source_MIGCTL {
		newReq := *req
		newReq.PeerAddress = r.config.Address
		newReq.PeerPort = int32(r.config.Port)
		newReq.Source = pb.Source_MIGRATORD

		// Fork the request to client and server
		response, err := r.PeerCreateMigrationJob(RPCPeer{
			Address: req.PeerAddress,
			Port:    int(req.PeerPort),
		}, newReq)

		if err != nil {
			return nil, err
		}

		return nil, nil
	} else {

	}
}

func (r *rpc) PeerCreateMigrationJob(peer RPCPeer, req pb.CreateMigrationJobRequest) (*pb.CreateMigrationJobResponse, error) {
	conn, err := grpc.Dial(peer.String(), grpc.WithInsecure())

	if err != nil {
		return nil, errors.New(fmt.Sprintf("call failed to peer ", peer.String()))
	}

	client := pb.NewMigratorServiceClient(conn)

	defer conn.Close()
	return client.CreateMigrationJob(context.Background(), &req)
}

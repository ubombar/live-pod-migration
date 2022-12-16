package main

import (
	"context"
	"flag"
	"fmt"

	pb "github.com/ubombar/live-pod-migration/pkg/migrator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var addrClient, addrServer string
	var portClient, portServer int

	// Default values are for debugging
	flag.StringVar(&addrClient, "addrc", "localhost", "Client address")
	flag.StringVar(&addrServer, "addrs", "localhost", "Server address")
	flag.IntVar(&portClient, "portc", 4545, "Client address")
	flag.IntVar(&portServer, "ports", 4546, "Server port")

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", addrClient, portClient), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	client := pb.NewMigratorServiceClient(conn)
	defer conn.Close()

	resp, err := client.CreateMigrationJob(context.Background(), &pb.CreateMigrationJobRequest{
		PeerAddress: addrServer,
		PeerPort:    int32(portServer),
		ContainerId: "59bca86ed5aa",
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("resp.Accepted: %v\n", resp)
}

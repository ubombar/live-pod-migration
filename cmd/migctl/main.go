package main

import (
	"context"
	"fmt"

	pb "github.com/ubombar/live-pod-migration/pkg/migrator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("migctl")

	addressClient := "localhost"
	portClient := 4545

	addressServer := "localhost"
	portServer := 4546

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", addressClient, portClient), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	client := pb.NewMigratorServiceClient(conn)
	defer conn.Close()

	resp, err := client.CreateMigrationJob(context.Background(), &pb.CreateMigrationJobRequest{
		PeerAddress: addressServer,
		PeerPort:    int32(portServer),
		ContainerId: "59bca86ed5aa",
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("resp.Accepted: %v\n", resp)
}

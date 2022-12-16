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

	conn, err := grpc.Dial("localhost:4545", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	client := pb.NewMigratorServiceClient(conn)
	defer conn.Close()

	resp, err := client.CreateMigrationJob(context.Background(), &pb.CreateMigrationJobRequest{
		PeerAddress: "localhost",
		PeerPort:    4545,
		ContainerId: "59bca86ed5aa",
	})

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("resp.Accepted: %v\n", resp)
}

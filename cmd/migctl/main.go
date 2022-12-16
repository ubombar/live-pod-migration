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

	resp, _ := client.MigrationJob(context.Background(), &pb.MigrationJobRequest{
		IncomingIp:    "localhost",
		FromMigratord: pb.MigrationJobSource_SOURCE_MIGCTL,
	})

	fmt.Printf("resp.Accepted: %v\n", resp.Accepted)

}

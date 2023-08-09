package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	pb "github.com/ubombar/live-pod-migration/pkg/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverPrivateKeyPath string
var serverUser string
var portClient int
var portServer int

var jobCmd = &cobra.Command{
	Use:   "job [OPTIONS] [source node address] [destination node address] [container name]",
	Short: "Create a new migration job",
	Long:  `Create a new migration job from the specitied plags.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("migctl job requires arguments [source node address] [destination node address] and [container name]")
			return
		}
		sourceNodeAddress := args[0]
		destinationNodeAddress := args[1]
		containerName := args[2]

		clientMigrator := fmt.Sprintf("%s:%d", sourceNodeAddress, portClient)
		conn, err := grpc.Dial(clientMigrator, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			fmt.Printf("Cannot reach the migratord client on %s\n", clientMigrator)
			return
		}

		client := pb.NewMigratorServiceClient(conn)
		defer conn.Close()

		// Do not include servers private key
		resp, err := client.CreateMigrationJob(context.Background(), &pb.CreateMigrationJobRequest{
			ContainerId:            containerName,
			ClientContainerRuntime: "podman",
			ServerContainerRuntime: "podman",
			ServerAddress:          destinationNodeAddress,
			ServerPort:             int32(portServer),
			ServerUser:             "ubombar",
			Method:                 "Basic",
		})

		if err != nil {
			fmt.Printf("Cannot create migration job on %s\n", clientMigrator)
			fmt.Printf("	Error: %v\n", err)
			return
		}

		fmt.Printf("%s\n", resp.MigrationId)
	},
}

func init() {
	jobCmd.Flags().IntVar(&portClient, "port-client", 9213, "Client's port")
	jobCmd.Flags().IntVar(&portServer, "port-server", 9213, "Server's port")

	jobCmd.Flags().StringVar(&serverPrivateKeyPath, "key", "~/.ssh/id_rsa", "id_rsa file of the server migratord.")
	jobCmd.Flags().StringVar(&serverUser, "user", "docker", "server's user for ssh connection.")

	rootCmd.AddCommand(jobCmd)
}

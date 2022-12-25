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

var jobCmd = &cobra.Command{
	Use:   "job [OPTIONS] [container_id]",
	Short: "Create a new migration job",
	Long:  `Create a new migration job from the specitied plags.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("migctl job requires one argument [container_id]")
			return
		}
		containerId := args[0]

		clientMigrator := fmt.Sprintf("%s:%d", rootConfig.addressClient, rootConfig.portClient)
		conn, err := grpc.Dial(clientMigrator, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			fmt.Printf("Cannot reach the migratord client on %s\n", clientMigrator)
			return
		}

		// privateKey, err := os.ReadFile(serverPrivateKeyPath)
		privateKey := ""

		if err != nil {
			fmt.Printf("Cannot load server's private key")
			return
		}

		client := pb.NewMigratorServiceClient(conn)
		defer conn.Close()

		resp, err := client.CreateMigrationJob(context.Background(), &pb.CreateMigrationJobRequest{
			ContainerId:            containerId,
			ClientContainerRuntime: "docker",
			ServerContainerRuntime: "docker",
			ServerAddress:          rootConfig.addressServer,
			ServerPort:             int32(rootConfig.portServer),
			ServerKey:              string(privateKey),
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
	jobCmd.Flags().StringVar(&serverPrivateKeyPath, "key", "~/.ssh/id_rsa", "id_rsa file of the server migratord.")
	jobCmd.Flags().StringVar(&serverUser, "user", "docker", "server's user for ssh connection.")
	rootCmd.AddCommand(jobCmd)
}

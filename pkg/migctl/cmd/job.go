package cmd

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
	pb "github.com/ubombar/live-pod-migration/pkg/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var portSource int32
var portDestination int32
var userSource string
var userDestination string
var method string

func getMigrationMethod(method string) (*pb.MigrationMethod, error) {
	value, ok := pb.MigrationMethod_value[method]

	if !ok {
		return nil, errors.New("unknown migration method name")
	}

	mig := pb.MigrationMethod(value)

	return &mig, nil
}

func createRandomMigrationId(length int) string {
	alphabet := []rune("0123456789abcdef")
	randomId := make([]rune, length)

	for i := 0; i < length; i++ {
		randomId[i] = alphabet[rand.Int31n(16)]
	}

	return string(randomId)
}

var jobCmd = &cobra.Command{
	Use:   "job [OPTIONS] [source node address] [destination node address] [container name or id]",
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
		randomMigrationId := createRandomMigrationId(32)

		migrationMethod, err := getMigrationMethod(method)

		if err != nil {
			fmt.Printf("Cannot create migration: %v\n", err)
		}

		clientMigrator := fmt.Sprintf("%s:%d", sourceNodeAddress, portSource)
		conn, err := grpc.Dial(clientMigrator, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			fmt.Printf("Cannot reach the migratord client on %s\n", clientMigrator)
			return
		}

		client := pb.NewMigratorServiceClient(conn)
		defer conn.Close()

		resp, err := client.CreateJob(context.Background(), &pb.CreateJobRequest{
			Source: &pb.IPAddress{
				Address: sourceNodeAddress,
				Port:    portSource,
			},
			Destination: &pb.IPAddress{
				Address: destinationNodeAddress,
				Port:    portDestination,
			},
			SourceUser:      userSource,
			DestinationUser: userDestination,
			MigrationMethod: *migrationMethod,
			RequestedBy:     pb.MigrationRole_controller,
			MigrationId:     randomMigrationId,
			ContainerName:   containerName,
		})

		if err != nil {
			fmt.Printf("Cannot create the migration %v\n", err)
			return
		}

		if resp.ErrorString != nil {
			fmt.Printf("Cannot create the migration %v\n", resp.ErrorString)
			return
		}

		fmt.Printf("%s\n", randomMigrationId)
	},
}

func init() {
	jobCmd.Flags().Int32Var(&portSource, "port-source", 9213, "Source's port")
	jobCmd.Flags().Int32Var(&portDestination, "port-destination", 9213, "Destination's port")

	jobCmd.Flags().StringVar(&userSource, "user-source", "vagrant", "Source machine's username")
	jobCmd.Flags().StringVar(&userDestination, "destination-source", "vagrant", "Destination machine's username")

	jobCmd.Flags().StringVar(&method, "method", "cold", "Migration method; cold, lazy, iterative")

	rootCmd.AddCommand(jobCmd)
}

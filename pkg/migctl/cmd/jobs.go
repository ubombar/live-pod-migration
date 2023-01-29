package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	pb "github.com/ubombar/live-pod-migration/pkg/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var nodePort int

var jobsCmd = &cobra.Command{
	Use:   "jobs [OPTIONS] [migratord address]",
	Short: "List the migration jobs in a host",
	Long:  `Display the migration jobs as a table in the command line.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("migctl job requires two arguments: [migratord address]")
			return
		}
		nodeAddress := args[0]

		migrator := fmt.Sprintf("%s:%d", nodeAddress, nodePort)
		conn, err := grpc.Dial(migrator, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			fmt.Printf("Cannot reach the migratord client on %s\n", migrator)
			return
		}

		client := pb.NewMigratorServiceClient(conn)
		defer conn.Close()

		// Do not include servers private key
		resp, err := client.GetMigrationJob(context.Background(), &pb.GetMigrationJobRequest{})

		if err != nil {
			fmt.Printf("Cannot get migration job on %s\n", migrator)
			fmt.Printf("	Error: %v\n", err)
			return
		}

		// resp := pb.GetMigrationJobResponse{
		// 	Jobs: []*pb.MigrationJob{
		// 		&pb.MigrationJob{
		// 			MigrationId:     "1324134123123",
		// 			CotninerId:      "323af34a21e09f",
		// 			MigrationStatus: "preparing",
		// 			CreationDate:    "10.10.2022",
		// 			Role:            "client",
		// 			MigrationMethod: "basic",
		// 		},
		// 	},
		// }

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", "Migration ID", "Container", "Stage", "Creation Date", "Role", "Method")
		for _, mj := range resp.Jobs {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", mj.MigrationId, mj.CotninerId, mj.MigrationStatus, mj.CreationDate, mj.Role, mj.MigrationMethod)
		}
		w.Flush()
	},
}

func init() {
	jobsCmd.Flags().IntVar(&nodePort, "port", 4545, "Port of the node we want to list migrations.")
	rootCmd.AddCommand(jobsCmd)
}

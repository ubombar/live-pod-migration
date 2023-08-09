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

var port int

var jobsCmd = &cobra.Command{
	Use:   "jobs [OPTIONS] [migratord address]",
	Short: "List the migration jobs in a host",
	Long:  `Display the migration jobs as a table in the command line.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("migctl job requires two arguments: [migratord address]")
			return
		}
		address := args[0]

		fullAddress := fmt.Sprintf("%s:%d", address, port)
		conn, err := grpc.Dial(fullAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			fmt.Printf("Cannot reach the migratord client on %s\n", fullAddress)
			return
		}

		client := pb.NewMigratorServiceClient(conn)
		defer conn.Close()

		resp, err := client.ListJobs(context.Background(), &pb.ListJobsRequest{})

		if err != nil {
			fmt.Printf("Cannot get migration job on %v: %v\n", fullAddress, err)
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", "Migration ID", "Running on Host", "Stage", "Creation Date", "Role", "Method")

		for _, job := range resp.GetJobs() {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", job.MigrationId, !job.Offline, job.MigrationState, job.CreationDate, job.Role, job.MigrationMethod)
		}

		w.Flush()
	},
}

func init() {
	jobsCmd.Flags().IntVar(&port, "port", 9213, "Port of the node we want to list migrations.")
	rootCmd.AddCommand(jobsCmd)
}

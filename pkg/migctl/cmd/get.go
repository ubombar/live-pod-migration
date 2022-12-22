package cmd

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"text/tabwriter"

// 	"github.com/spf13/cobra"
// 	pb "github.com/ubombar/live-pod-migration/pkg/migrator/"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// var getCmd = &cobra.Command{
// 	Use:   "get",
// 	Short: "Get information from nodes and the status",
// 	Long:  `Get migration jobs from both node and compare.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		if len(args) < 1 {
// 			fmt.Println("migctl job requires one argument [container_id]")
// 			return
// 		}
// 		migrationid := args[0]

// 		clientMigrator := fmt.Sprintf("%s:%d", rootConfig.addressClient, rootConfig.portClient)
// 		serverMigrator := fmt.Sprintf("%s:%d", rootConfig.addressServer, rootConfig.portServer)

// 		connClient, err := grpc.Dial(clientMigrator, grpc.WithTransportCredentials(insecure.NewCredentials()))

// 		if err != nil {
// 			fmt.Printf("Cannot reach the migratord client on %s\n", clientMigrator)
// 			return
// 		}

// 		connServer, err := grpc.Dial(serverMigrator, grpc.WithTransportCredentials(insecure.NewCredentials()))

// 		if err != nil {
// 			fmt.Printf("Cannot reach the migratord client on %s\n", serverMigrator)
// 			return
// 		}

// 		clientServer := pb.NewMigratorServiceClient(connServer)
// 		clientClient := pb.NewMigratorServiceClient(connClient)

// 		defer connClient.Close()
// 		defer connServer.Close()

// 		respClient, err := clientClient.GetMigrationStatus(context.Background(), &pb.GetMigrationStatusRequest{
// 			MigrationId: migrationid,
// 		})

// 		if err != nil {
// 			fmt.Printf("Cannot find migration job on %s\n", clientMigrator)
// 			fmt.Printf("	Error: %v\n", err)
// 			return
// 		}

// 		respServer, err := clientServer.GetMigrationStatus(context.Background(), &pb.GetMigrationStatusRequest{
// 			MigrationId: migrationid,
// 		})

// 		if err != nil {
// 			fmt.Printf("Cannot find migration job on %s\n", clientMigrator)
// 			fmt.Printf("	Error: %v\n", err)
// 			return
// 		}

// 		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
// 		fmt.Fprintf(w, "%v\t%v\t%v\n", "Role", respServer.MigrationRole, respClient.MigrationRole)
// 		fmt.Fprintf(w, "%v\t%v\t%v\n", "Client ContainerId", respServer.ClientContainerId, respClient.ClientContainerId)
// 		fmt.Fprintf(w, "%v\t%v\t%v\n", "Server ContainerId", respServer.ServerContainerId, respClient.ServerContainerId)
// 		fmt.Fprintf(w, "%v\t%v\t%v\n", "Method", respServer.MigrationMethod, respClient.MigrationMethod)
// 		fmt.Fprintf(w, "%v\t%v\t%v\n", "Running", respServer.Running, respClient.Running)
// 		fmt.Fprintf(w, "%v\t%v\t%v\n", "Status", respServer.MigrationStatus, respClient.MigrationStatus)
// 		w.Flush()
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(getCmd)
// }

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type MigCTLConfig struct {
	addressClient string
	addressServer string

	portClient int
	portServer int
}

var rootConfig = MigCTLConfig{}

var rootCmd = &cobra.Command{
	Short: "migctl is a command line utility for triggering live migrations.",
	Long: `migctl is a command line utility for triggering live migrations.
The nodes should contain a utility named 'migratord'.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use migctl --help to get help.")
	},
	Hidden: true,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&rootConfig.addressClient, "address-client", "localhost", "Client's IP address")
	rootCmd.PersistentFlags().StringVar(&rootConfig.addressServer, "address-server", "localhost", "Server's IP address")
	rootCmd.PersistentFlags().IntVar(&rootConfig.portClient, "port-client", 4545, "Client's port")
	rootCmd.PersistentFlags().IntVar(&rootConfig.portServer, "port-server", 4545, "Server's port")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

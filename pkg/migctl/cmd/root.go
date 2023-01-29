package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "migctl is a command line utility for triggering live migrations.",
	Long: `migctl is a command line utility for triggering live migrations.
The nodes should contain a utility named 'migratord'.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use migctl --help to get help.")
	},
	Hidden: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of migctl",
	Long:  `All software has versions. This is migctl's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migctl version v1alpha1 1.01")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

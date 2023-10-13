package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cia",
	Short: "Container Image Analyzer - Analyze Docker container images",
	Run: func(cmd *cobra.Command, args []string) {
		// Implement your logic here
		fmt.Println("Analyzing container image...")
	},
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

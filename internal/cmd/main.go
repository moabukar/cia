package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ciaCmd = &cobra.Command{
	Use:   "cia",
	Short: "Container Image Analyzer - Analyze Docker container images",
	Run: func(cmd *cobra.Command, args []string) {
		// Implement your logic here
		fmt.Println("Analyzing container image...")
	},
}

func Execute() error {
	return ciaCmd.Execute()
}

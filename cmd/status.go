// Package cmd contains descriptions and handlers for Pixie Clock Adapter CLI.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// status represents the status command.
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Prints application status",
	Run: func(_ *cobra.Command, _ []string) {
		d := getDaemon()
		if d.Running() {
			fmt.Println("Running")
		} else {
			fmt.Println("Not running")
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

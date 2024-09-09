// Package cmd contains descriptions and handlers for Pixie Clock Adapter CLI.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command.
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the background application",
	Run: func(_ *cobra.Command, _ []string) {
		d := getDaemon()
		if d.Running() {
			err := d.Stop()
			if err != nil {
				fmt.Println("Error while stopping:", err.Error())
			} else {
				fmt.Println("Daemon is stopped.")
			}
		} else {
			fmt.Println("Daemon is not running.")
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

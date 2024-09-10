// Package cmd contains descriptions and handlers for Pixie Clock Adapter CLI.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// startCmd represents the start command.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the application in the background",
	Run: func(_ *cobra.Command, _ []string) {
		s := getApplication()
		d := getDaemon()
		if d.Running() {
			fmt.Println("Application is already running in background.")
			os.Exit(1)
		}
		child, err := d.Context.Reborn()
		if err != nil {
			fmt.Println("Error while starting daemon:", err.Error())
		}
		if child == nil {
			defer d.Context.Release() //nolint:errcheck
			err = s.Start()
			if err != nil {
				fmt.Println("Error while starting application:", err.Error())
				os.Exit(1)
			}
		} else {
			fmt.Println("Daemon is started.")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

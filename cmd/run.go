// Package cmd contains descriptions and handlers for Pixie Clock Adapter CLI.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// runCmd represents the run command.
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the application in the current process",
	Run: func(_ *cobra.Command, _ []string) {
		s := getApplication()
		d := getDaemon()
		if d.Running() {
			fmt.Println("Application is running in background. Stop it, before run")
			os.Exit(1)
		}
		panic(s.Start())
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

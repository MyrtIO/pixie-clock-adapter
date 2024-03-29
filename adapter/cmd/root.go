// Package cmd contains descriptions and handlers for Pixie Clock Adapter CLI.
package cmd

import (
	"os"
	"pixie_adapter/internal/app"
	"pixie_adapter/pkg/process"

	"github.com/spf13/cobra"
)

// AppName represents app name.
const AppName = "pixie-adapter"

// PackageName represents app package name.
const PackageName = "co.myrt.pixie_adapter"

// Version represents current app version.
var Version = "snapshot"

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     AppName,
	Version: Version,
	Short:   "An application that provides an HTTP REST API for Pixie Clock",
}

// Execute is the main CLI entrypoint.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var baudRate int
var portPath string
var httpPort int

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&portPath,
		"port", "p",
		"",
		"Port (starting with /dev/)")
	rootCmd.PersistentFlags().IntVarP(
		&baudRate,
		"baud", "b",
		28800,
		"Set baud/speed")
	rootCmd.PersistentFlags().IntVarP(
		&httpPort,
		"http-port", "t",
		17085,
		"HTTP port where API will be available")
}

func getDaemon() process.Daemon {
	return process.NewDaemon(PackageName)
}

func getService() *app.Application {
	app := app.New(portPath, baudRate)
	app.SetPort(httpPort)
	return app
}

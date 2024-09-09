// Package cmd contains descriptions and handlers for Pixie Clock Adapter CLI.
package cmd

import (
	"fmt"
	"os"
	"pixie_adapter/internal/app"
	"pixie_adapter/internal/config"
	"pixie_adapter/pkg/process"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     config.AppName,
	Version: config.Version,
	Short:   "An application that provides an HTTP REST API for Pixie Clock",
}

// Execute is the main CLI entrypoint.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var configPath string

func init() {
	homeDir := os.Getenv("HOME")

	rootCmd.PersistentFlags().StringVarP(
		&configPath,
		"config", "c",
		fmt.Sprintf("%s/.config/%s/config.yaml", homeDir, config.AppName),
		"The path to the configuration file",
	)
}

func getDaemon() process.Daemon {
	return process.NewDaemon(config.PackageName)
}

func getService() *app.Application {
	app := app.New(configPath)
	return app
}

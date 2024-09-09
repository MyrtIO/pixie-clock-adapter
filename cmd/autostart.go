// Package cmd contains descriptions and handlers for vpn-dns CLI.
package cmd

import (
	"fmt"
	"os"
	"pixie_adapter/internal/config"

	"github.com/mishamyrt/go-lunch"
	"github.com/spf13/cobra"
)

func getAgent() *lunch.Agent {
	binPath, err := os.Executable()
	if err != nil {
		fmt.Println("Executable path is not found:", err.Error())
		os.Exit(1)
	}
	path, _ := lunch.UserPath(config.PackageName)
	agent := lunch.Agent{
		PackageName: config.PackageName,
		Command:     binPath + " start",
		Path:        path,
	}
	return &agent
}

// autostartEnableCmd represents the `autostart enable` command.
var autostartEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enables automatic startup",
	Run: func(_ *cobra.Command, _ []string) {
		agent := getAgent()
		exists, _ := agent.Exists()
		if exists {
			fmt.Println("Autostart is already enabled.")
			os.Exit(1)
		}
		err := agent.Write()
		if err != nil {
			fmt.Println("Can't write login item:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Autostart is enabled.")
	},
}

// autostartDisableCmd represents the `autostart disable` command.
var autostartDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disables automatic startup",
	Run: func(_ *cobra.Command, _ []string) {
		agent := getAgent()
		exists, _ := agent.Exists()
		if !exists {
			fmt.Println("Autostart is not enabled.")
			os.Exit(1)
		}
		err := agent.Remove()
		if err != nil {
			fmt.Println("Can't remove login item:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Autostart is disabled.")
	},
}

// autostartCmd represents the autostart command.
var autostartCmd = &cobra.Command{
	Use:   "autostart",
	Short: "Controls the automatic start-up of the application",
	Run: func(_ *cobra.Command, _ []string) {
		agent := getAgent()
		exists, _ := agent.Exists()
		if exists {
			fmt.Println("Autostart is enabled.")
			fmt.Println("To disable, run: pixie-adapter autostart disable")
		} else {
			fmt.Println("Autostart is not enabled.")
			fmt.Println("To enable, run: pixie-adapter autostart enable")
		}
	},
}

func init() {
	autostartCmd.AddCommand(autostartEnableCmd)
	autostartCmd.AddCommand(autostartDisableCmd)
	rootCmd.AddCommand(autostartCmd)
}

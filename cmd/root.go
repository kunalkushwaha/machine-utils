package cmd

import (
	"fmt"
	"os"

	"github.com/docker/machine/commands/mcndirs"
	"github.com/spf13/cobra"
)

// needs export=docker-machine-vm-base-path
var machinePath string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "machine-utils",
	Short: "provides few missing functions of docker-machine",
	Long: `Provides few extra functionality from docker-machine.

This tool complementry tool to docker-machine
It uses same environment variables of docker-machine,
So without docker-machine it may not be usable`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	machinePath = os.Getenv("docker_machine_dir")
	if machinePath == "" {
		machinePath = mcndirs.GetBaseDir()
	}
	fmt.Println(machinePath)
}

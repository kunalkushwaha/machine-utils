package cmd

import (
	"fmt"

	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/log"
	"github.com/spf13/cobra"
)

// provisionCmd represents the provision command
var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Provision docker-host",
	Long:  `Provision docker-host created by docker-machine`,

	Run: ProvisionCmd,
}

func init() {
	RootCmd.AddCommand(provisionCmd)

	provisionCmd.Flags().StringP("machine", "m", "default", "machine to be provisioned")
	provisionCmd.Flags().StringP("file", "f", "machine-config.yml", "Config file for provisioner")

}

// ProvisionCmd provision's as per config file.
func ProvisionCmd(cmd *cobra.Command, args []string) {
	machine, _ := cmd.Flags().GetString("machine")
	file, _ := cmd.Flags().GetString("file")
	log.Info("Machine: ", machine)
	log.Info("File: ", file)

	client := libmachine.NewClient(machinePath, machinePath)
	defer client.Close()

	h, err := client.Load(machine)
	if err != nil {
		log.Error(err)
		return
	}

	if err := h.Start(); err != nil {
		log.Error(err)
	}

	out, err := h.RunSSHCommand("df -h")
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Printf("Results of your disk space query:\n%s\n", out)

}

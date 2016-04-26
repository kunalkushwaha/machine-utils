package cmd

import (
	//	"fmt"
	"os"

	"github.com/docker/machine/libmachine/drivers/plugin/localbinary"
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

	provisionCmd.Flags().StringP("file", "f", "machine-config.yml", "Config file for provisioner")

}

func ProvisionCmd(cmd *cobra.Command, args []string) {
}

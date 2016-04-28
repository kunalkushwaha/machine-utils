package cmd

import "github.com/spf13/cobra"

// provisionCmd represents the provision command
var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Provision docker-host",
	Long:  `Provision docker-host created by docker-machine`,

	Run: ProvisionCmd,
}

func init() {
	RootCmd.AddCommand(provisionCmd)

	provisionCmd.PersistentFlags().StringP("machine", "m", "default", "machine to be provisioned")
	provisionCmd.PersistentFlags().StringP("file", "f", "machine-config.yml", "Config file for provisioner")

}

// ProvisionCmd provision's as per config file.
func ProvisionCmd(cmd *cobra.Command, args []string) {
	// TODO:
	// Parse yml file and provision accordingly.
}

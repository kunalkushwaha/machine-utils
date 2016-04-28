package cmd

import (
	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/log"
	"github.com/spf13/cobra"
)

// storageCmd represents the storage command
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Configures and prepares storage for docker containers.",
	Long:  `Install and configure storage for containers using storage plugin.`,
	Run:   StorageCmd,
}

func init() {
	provisionCmd.AddCommand(storageCmd)
	// TODO:
	// Options ISCSI, NFS, EC3
}

// StorageCmd configures storage for docker-host.
func StorageCmd(cmd *cobra.Command, args []string) {
	machine, _ := cmd.Flags().GetString("machine")
	file, _ := cmd.Flags().GetString("file")
	log.Debug("Machine: ", machine)
	log.Debug("File: ", file)

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

	/*
		 TODO:
		 - Every docker plugin works on socket.
			- This can be achived by running them in docker-container.
			- But keeping the official/Standard docker-image important.
			- Check comptaiblity with docker swarm too!

			- Intial storage provider should
				 - NFS
				 - ISCSI
				 - EC3
	*/

}

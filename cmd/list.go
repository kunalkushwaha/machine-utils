package cmd

import (
	"fmt"
	"os"

	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/log"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available machines",
	Long:  `List available machines created by docker-machine`,
	Run:   ListCmd,
}

func init() {
	RootCmd.AddCommand(listCmd)
}

// ListCmd list all the machines available
func ListCmd(cmd *cobra.Command, args []string) {
	// needs export=docker-machine-vm-base-path
	machinePath := os.Getenv("docker_machine_dir")
	client := libmachine.NewClient(machinePath, machinePath)
	defer client.Close()

	hosts, err := client.List()
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println("Docker Host")
	fmt.Println("-----------")
	for _, host := range hosts {
		fmt.Println(host)
	}
}

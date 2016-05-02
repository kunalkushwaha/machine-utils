package cmd

import (
	"fmt"
	"strings"

	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/log"
	"github.com/spf13/cobra"
)

var cAdvisorTemplateCmd = `sudo docker run \
  --volume=/:/rootfs:ro \
  --volume=/var/run:/var/run:rw \
  --volume=/sys:/sys:ro \
  --volume=/var/lib/docker/:/var/lib/docker:ro \
  --publish=8080:8080 \
  --detach=true \
  --name=cadvisor \
	--restart=always \
  google/cadvisor:latest`

// cadvisorCmd represents the cadvisor command
var cadvisorCmd = &cobra.Command{
	Use:   "cadvisor",
	Short: "Configures cAdvisor on docker host.",
	Long:  `Configures cAdvisor on docker host for container monitoring.`,
	Run:   CAdvisorCmd,
}

func init() {
	provisionCmd.AddCommand(cadvisorCmd)
}

// CAdvisorCmd install and configures the cAdvisor on docker host
func CAdvisorCmd(cmd *cobra.Command, args []string) {

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

	_, err = h.RunSSHCommand("docker ps | grep cadvisor")
	if err != nil {
		_, err = h.RunSSHCommand(cAdvisorTemplateCmd)
		if err != nil {
			log.Error(err)
		}
	}

	tcpURL, _ := h.URL()
	url := strings.Split(tcpURL, ":")

	fmt.Printf("cAdvisor is successfully  setup\nYou can access it at http:%s:8080 \n", url[1])

}

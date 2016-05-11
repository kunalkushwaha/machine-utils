package cmd

import (
	"fmt"

	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type osdConfig struct {
	Osd `yaml:"osd"`
}

// Osd struct
type Osd struct {
	Drivers `yaml:"drivers"`
}

// Drivers struct
type Drivers struct {
	Nfs `yaml:"nfs"`
	//	Aws `yaml:"aws"`
}

// Nfs struct
type Nfs struct {
	Server string `yaml:"server"`
	Path   string `yaml:"path"`
}

// Aws struct
type Aws struct {
	AwsAccessKeyID     string `yaml:"aws_access_key_id"`
	AWSSecretAccessKey string `yaml:"aws_secret_access_key"`
}

var osdCommand = `
docker run \
		--privileged \
		-d \
		-v /home/docker:/etc \
		-v /run/docker/plugins:/run/docker/plugins \
		-v /var/lib/openstorage:/var/lib/openstorage \
		-v /var/lib/osd/:/var/lib/osd/ \
		-p 2345:2345 \
		--restart always \
		openstorage/osd -d -f /etc/config.yaml
`

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
	storageCmd.Flags().String("nfs-ip", "", "IP Address of nfs server")
	storageCmd.Flags().String("nfs-path", "", "Path of nfs share")
}

// StorageCmd configures storage for docker-host.
func StorageCmd(cmd *cobra.Command, args []string) {
	machine, _ := cmd.Flags().GetString("machine")
	file, _ := cmd.Flags().GetString("file")
	nfsIP, _ := cmd.Flags().GetString("nfs-ip")
	nfsPath, _ := cmd.Flags().GetString("nfs-path")

	log.Debug("Machine: ", machine)
	log.Debug("File: ", file)

	if len(nfsIP) > 0 || len(nfsPath) > 0 {
		if len(nfsPath) == 0 || len(nfsPath) == 0 {
			fmt.Println("Please provide 'nfs-ip' and 'nfs-path'")
			return
		}
	}

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

	config := osdConfig{}
	config.Nfs.Server = nfsIP
	config.Nfs.Path = nfsPath
	configData, err := yaml.Marshal(&config)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	cmdEcho := "echo \"" + string(configData) + "\" > ~/config.yaml"
	_, err = h.RunSSHCommand(cmdEcho)
	if err != nil {
		fmt.Println("Something wrong at config file creation!")
		return
	}

	_, err = h.RunSSHCommand(osdCommand)
	if err != nil {
		fmt.Println(err)
		return
	}

}

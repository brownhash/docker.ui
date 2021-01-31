package docker

import (
	"github.com/docker/docker/client"
	"github.com/sharma1612harshit/gomuf/logs"
)

// Client - return docker sdk client
func Client() (*client.Client, error) {
	logs.Debug("Initiating docker client")

	cli, err := client.NewEnvClient()
	if err != nil {
		logs.Debug(err)
	}

	return cli, err
}

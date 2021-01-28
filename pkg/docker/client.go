package docker

import (
	"github.com/docker/docker/client"
	"github.com/sharma1612harshit/docker.ui/pkg/logger"
)

// Client - return docker sdk client
func Client() (*client.Client, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		logger.Debug(err)
	}

	return cli, err
}

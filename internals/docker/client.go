package docker

import "github.com/docker/docker/client"

func Client() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	return cli, err
}

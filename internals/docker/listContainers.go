package docker

import (
	"context"
	"github.com/docker/docker/api/types"
)

func ListContainers() ([]types.Container, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		return nil, err
	}

	list, err := cli.ContainerList(ctx, types.ContainerListOptions{})

	return list, err
}

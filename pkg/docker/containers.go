package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
)

// List local docker containers
func GetContainers(quiet bool, size bool, all bool, latest bool, since string, before string, limit int, filters filters.Args) ([]types.Container, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		return nil, err
	}

	list, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Quiet:   quiet, 		// false
		Size:    size, 			// false
		All:     all, 			// true
		Latest:  latest, 		// false
		Since:   since, 		// ""
		Before:  before, 		// ""
		Limit:   limit, 		// 0
		Filters: filters, 		// {}
	})

	return list, err
}

func LaunchContainer(containerConfig *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}

	launchResponse, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, networkingConfig, containerName)

	return launchResponse, err
}

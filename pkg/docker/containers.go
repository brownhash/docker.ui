package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/sharma1612harshit/docker.ui/pkg/logger"
)

// GetContainers - List local docker containers
func GetContainers(quiet, size, all, latest bool, since, before string, limit int) ([]types.Container, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		logger.Debug(err)
		return nil, err
	}

	list, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Quiet:   quiet, 			// false
		Size:    size, 				// false
		All:     all, 				// true
		Latest:  latest, 			// false
		Since:   since, 			// ""
		Before:  before, 			// ""
		Limit:   limit, 			// 0
		Filters: filters.Args{}, 	// {}
	})
	
	logger.Debug(err)

	return list, err
}

// LaunchContainer - launch a conatiner
func LaunchContainer(containerConfig *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		logger.Debug(err)
		return container.ContainerCreateCreatedBody{}, err
	}

	launchResponse, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, networkingConfig, containerName)
	
	logger.Debug(err)

	return launchResponse, err
}

package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/sharma1612harshit/gomuf/logs"
)

// GetContainers - List local docker containers
func GetContainers(quiet, size, all, latest bool, since, before string, limit int) ([]types.Container, error) {
	logs.Debug(fmt.Sprintf("Initiating GetContainers with \nquiet: %v, \nsize: %v, \nall: %v, \nlatest: %v, \nsince: %v, \nbefore: %v, \nlimit: %v", quiet, size, all, latest, since, before, limit))

	logs.Debug("Initiating background context")
	ctx := context.Background()

	logs.Debug("Initiating docker sdk client")
	cli, err := Client()

	if err != nil {
		logs.Debug(err)
		return nil, err
	}

	logs.Debug("Fetching containers")
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
	
	if err != nil {
		logs.Debug(err)
	}

	return list, err
}

// LaunchContainer - launch a conatiner
func LaunchContainer(containerConfig *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	logs.Debug(fmt.Sprintf("Initiating LaunchContainer with \ncontainerConfig: %v, \nhostConfig: %v, \nnetworkingConfig: %v, \ncontainerName: %v", containerConfig, hostConfig, networkingConfig, containerName))

	logs.Debug("Initiating background context")
	ctx := context.Background()

	logs.Debug("Initiating docker sdk client")
	cli, err := Client()

	if err != nil {
		logs.Debug(err)
		return container.ContainerCreateCreatedBody{}, err
	}

	logs.Debug("launching container and generating response")
	launchResponse, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, networkingConfig, containerName)
	
	if err != nil {
		logs.Debug(err)
	}

	return launchResponse, err
}

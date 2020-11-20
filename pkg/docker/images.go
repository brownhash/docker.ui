package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

// List local docker images
func GetImages() ([]types.ImageSummary, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		return nil, err
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{
		All:     true,
		Filters: filters.Args{},
	})

	return images, err
}

package docker

import (
	"context"
	"github.com/docker/docker/api/types"
)

func ListImages() ([]types.ImageSummary, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		return nil, err
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{})

	return images, err
}

package docker

import (
	"context"
	"github.com/docker/docker/api/types"
)

func DeleteImage(imageId string) ([]types.ImageDeleteResponseItem, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		return nil, err
	}

	dResponse, err := cli.ImageRemove(ctx, imageId, types.ImageRemoveOptions{
		Force:         false,
		PruneChildren: false,
	})

	return dResponse, err
}

package docker

import (
	"context"
	"github.com/docker/docker/api/types"
)

func DeleteImage(imageId string, force bool, pruneChildren bool) ([]types.ImageDeleteResponseItem, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		return nil, err
	}

	dResponse, err := cli.ImageRemove(ctx, imageId, types.ImageRemoveOptions{
		Force:         force,
		PruneChildren: pruneChildren,
	})

	return dResponse, err
}

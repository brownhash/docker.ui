package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"io"
	"os"
)

// List local docker images
func GetImages(all bool, filter filters.Args) ([]types.ImageSummary, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		return nil, err
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{
		All:     all,				// true
		Filters: filter, 	// {}
	})

	return images, err
}

// Pull specified docker image
func PullImage(imageRef string, all bool, username string, password string) error {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		return err
	}

	authEncode, err := json.Marshal(types.AuthConfig{
		Username: username,
		Password: password,
	})

	if err != nil {
		return err
	}

	authStr := base64.URLEncoding.EncodeToString(authEncode)

	reader, err := cli.ImagePull(ctx, imageRef, types.ImagePullOptions{
		All:           all,
		RegistryAuth:  authStr,
		PrivilegeFunc: nil,
	})

	if err != nil {
		return err
	}
	_, err = io.Copy(os.Stdout, reader)

	return err
}

// Delete specified docker image
func DeleteImage(imageId string, force bool, pruneChildren bool) ([]types.ImageDelete, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		return nil, err
	}

	dResponse, err := cli.ImageRemove(ctx, imageId, types.ImageRemoveOptions{
		Force:         force,			// false
		PruneChildren: pruneChildren,	// false
	})

	return dResponse, err
}

package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/sharma1612harshit/docker.ui/pkg/logger"
	"io"
	"os"
)

// GetImages - List local docker images
func GetImages(all bool, filter filters.Args) ([]types.ImageSummary, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		logger.Debug(err)
		return nil, err
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{
		All:     all,				// true
		Filters: filter, 	// {}
	})

	if err != nil {
		logger.Debug(err)
	}

	return images, err
}

// PullImage - Pull specified docker image
func PullImage(imageRef string, all bool, username, password string) error {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		logger.Debug(err)
		return err
	}

	authEncode, err := json.Marshal(types.AuthConfig{
		Username: username,
		Password: password,
	})

	if err != nil {
		logger.Debug(err)
		return err
	}

	authStr := base64.URLEncoding.EncodeToString(authEncode)

	reader, err := cli.ImagePull(ctx, imageRef, types.ImagePullOptions{
		All:           all,
		RegistryAuth:  authStr,
		PrivilegeFunc: nil,
	})

	if err != nil {
		logger.Debug(err)
		return err
	}
	_, err = io.Copy(os.Stdout, reader)

	if err != nil {
		logger.Debug(err)
	}

	return err
}

// DeleteImage - Delete specified docker image
func DeleteImage(imageID string, force, pruneChildren bool) ([]types.ImageDelete, error) {
	ctx := context.Background()

	cli, err := Client()

	if err != nil {
		logger.Debug(err)
		return nil, err
	}

	dResponse, err := cli.ImageRemove(ctx, imageID, types.ImageRemoveOptions{
		Force:         force,			// false
		PruneChildren: pruneChildren,	// false
	})

	if err != nil {
		logger.Debug(err)
	}

	return dResponse, err
}

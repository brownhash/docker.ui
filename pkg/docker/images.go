package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/sharma1612harshit/docker.ui/pkg/logger"
	"io"
	"os"
)

// GetImages - List local docker images
func GetImages(all bool, filter filters.Args) ([]types.ImageSummary, error) {
	logger.Debug(fmt.Sprintf("Initiating GetImages with \nall: %v, \nfilter: %v", all, filter))

	logger.Debug("Initiating background context")
	ctx := context.Background()

	logger.Debug("Initiating docker sdk client")
	cli, err := Client()

	if err != nil {
		logger.Debug(err)
		return nil, err
	}

	logger.Debug("Fetching images and generating response")
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
	logger.Debug(fmt.Sprintf("Initiating PullImage with \nimageRef: %v, \nall: %v, \nusername: %v, \npassword: %v", imageRef, all, username, password))

	logger.Debug("Initiating background context")
	ctx := context.Background()

	logger.Debug("Initiating docker sdk client")
	cli, err := Client()

	if err != nil {
		logger.Debug(err)
		return err
	}

	logger.Debug("generating authConfig")
	authEncode, err := json.Marshal(types.AuthConfig{
		Username: username,
		Password: password,
	})

	if err != nil {
		logger.Debug(err)
		return err
	}

	logger.Debug("encoding authConfig")
	authStr := base64.URLEncoding.EncodeToString(authEncode)

	logger.Debug("pulling image and generating response")
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
	logger.Debug(fmt.Sprintf("Initiating DeleteImage with \nimageID: %v, \nforce: %v, \npruneChildren: %v", imageID, force, pruneChildren))

	logger.Debug("Initiating background context")
	ctx := context.Background()

	logger.Debug("Initiating docker sdk client")
	cli, err := Client()

	if err != nil {
		logger.Debug(err)
		return nil, err
	}

	logger.Debug("Deleting image and generating response")
	dResponse, err := cli.ImageRemove(ctx, imageID, types.ImageRemoveOptions{
		Force:         force,			// false
		PruneChildren: pruneChildren,	// false
	})

	if err != nil {
		logger.Debug(err)
	}

	return dResponse, err
}

package docker

import (
	"github.com/docker/docker/api/types/filters"
	"github.com/sharma1612harshit/docker.ui/pkg/docker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	// test get images
	_, iErr := docker.GetImages(true, filters.Args{})
	assert.Empty(t, iErr, "Error occurred while running pkg/docker/GetImages. Error %v.", iErr)

	// test delete image
	_, idErr := docker.DeleteImage("test_image_id", false, false)
	assert.Empty(t, idErr, "Error occurred while running pkg/docker/DeleteImage. Error %v.", idErr)

	// test get containers
	_, cErr := docker.GetContainers(false, false, true, false, "", "", 0, filters.Args{})
	assert.Empty(t, cErr, "Error occurred while running pkg/docker/GetContainers. Error %v.", cErr)
}
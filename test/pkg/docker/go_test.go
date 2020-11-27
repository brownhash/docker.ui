package docker

import (
	"github.com/sharma1612harshit/docker.ui/pkg/docker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	// test pull image
	ipErr := docker.PullImage("docker.io/library/golang", false, "", "")
	assert.Empty(t, ipErr, "Error occurred while running pkg/docker/PullImage. Error %v.", ipErr)

	// test get images
	_, iErr := docker.GetImages(true)
	assert.Empty(t, iErr, "Error occurred while running pkg/docker/GetImages. Error %v.", iErr)

	// test delete image
	_, idErr := docker.DeleteImage("golang", true, true)
	assert.Empty(t, idErr, "Error occurred while running pkg/docker/DeleteImage. Error %v.", idErr)

	// test get containers
	_, cErr := docker.GetContainers(false, false, true, false, "", "", 0)
	assert.Empty(t, cErr, "Error occurred while running pkg/docker/GetContainers. Error %v.", cErr)
}

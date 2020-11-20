package docker

import (
	"context"
	"github.com/docker/docker/client"
)

func Client() int {
	ctx := context.Background()
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}

	return client
}

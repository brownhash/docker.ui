
package docker

import "github.com/docker/docker/client"

func Client() (*client.Client, error) {
	cli, err := client.NewEnvClient()

	return cli, err
}
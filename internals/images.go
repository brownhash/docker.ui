package internals

import (
	"./docker"

	"./modals"
)

func Images() ([]modals.Image, error) {
	images, err := docker.ListImages()

	if err != nil {
		return nil, err
	}

	var imageList []modals.Image

	for _, image := range images {
		tmp := modals.Image{
			Id:         image.ID,
			RepoTag:    image.RepoTags,
			Size:       image.Size,
			RepoDigest: image.RepoDigests,
			CreatedAt:  image.Created,
			Labels:     image.Labels,
			Containers: image.Containers,
		}
		imageList = append(imageList, tmp)
	}

	return imageList, nil
}

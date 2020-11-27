package docker

import (
	"github.com/sharma1612harshit/docker.ui/pkg/docker"
	"log"
)

// return images data as json map
func GetImages() ([]ImageResponse, error) {
	images, err := docker.GetImages(true)

	var imageList = make([]ImageResponse, 0)

	if err != nil {
		log.Fatal(err)
		return imageList, err
	}

	for _, data := range images {
		imageList = append(imageList, ImageResponse{
			ID:         data.ID,
			Created:    data.Created,
			Containers: data.Containers,
			Labels:     data.Labels,
			RepoDigest: data.RepoDigests,
			RepoTags:   data.RepoTags,
			Size:       data.Size,
			ParentId:   data.ParentID,
		})
	}

	return imageList, err
}

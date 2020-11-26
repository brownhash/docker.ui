package docker

import (
	"github.com/docker/docker/api/types/filters"
	"github.com/sharma1612harshit/docker.ui/pkg/docker"
	"log"
)

// return images data as json map
func GetImages() map[string][]ImageResponse {
	images, err := docker.GetImages(true, filters.Args{})

	if err != nil {
		log.Fatal(err)
	}

	var imageList []ImageResponse

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

	response := map[string][]ImageResponse{
		"images": imageList,
	}

	return response
}

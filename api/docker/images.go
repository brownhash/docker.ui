package docker

import (
	"github.com/sharma1612harshit/docker.ui/pkg/docker"
	"log"
)

// return images data as json map
func GetImages() (map[string][]ImageResponse, error) {
	images, err := docker.GetImages(true)

	var response = map[string][]ImageResponse{}

	if err != nil {
		log.Fatal(err)
		return response, err
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

	response["images"] = imageList

	return response, err
}

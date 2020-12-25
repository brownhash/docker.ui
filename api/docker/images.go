package docker

import (
	"encoding/json"
	"github.com/docker/docker/api/types/filters"
	"github.com/sharma1612harshit/docker.ui/pkg/docker"
	"log"
)

// return images data as json map
func GetImages(all string, filter string) ([]ImageResponse, error) {
	allImages := false
	searchFilters := map[string]map[string]bool{}

	if all == "true" {
		allImages = true
	}

	json.Unmarshal([]byte(filter), &searchFilters)

	log.Print(searchFilters, filter)

	images, err := docker.GetImages(allImages, filters.Args{})

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

package docker

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/docker/docker/api/types/filters"
	"github.com/sharma1612harshit/docker.ui/pkg/docker"
)

// return images data as json map
func GetImages(all, filter string) ([]ImageResponse, error) {
	allImages := false
	searchFilters := map[string]map[string]bool{}

	if all == "true" {
		allImages = true
	}

	err := json.Unmarshal([]byte(filter), &searchFilters)

	if err != nil {
		log.Print(err)
	}

	// add filters to filter var
	var searchFilter = filters.NewArgs()
	
	for key, value := range(searchFilters) {
		for name, boolean := range(value) {
			searchFilter.Add(key, fmt.Sprintf("{\"%s\":%v}",name, boolean))
		}
	}

	images, err := docker.GetImages(allImages, searchFilter)

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

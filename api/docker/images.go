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

	json.Unmarshal([]byte(filter), &searchFilters)

	// add filters to filter var
	var searchFilter = filters.NewArgs()
	
	for key, value := range(searchFilters) {
		if key == "label" || key == "created" {
			for name, boolean := range(value) {
				searchFilter.Add(key, fmt.Sprintf("{\"%s\":%v}",name, boolean))
			}
		} else {
			log.Print("Invalid filter passed: ", key)
		}
	}

	images, err := docker.GetImages(allImages, searchFilter)

	var imageList = make([]ImageResponse, 0)

	if err != nil {
		log.Print(err)
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

func PullImage(all, reference, username, password string) (string, error) {
	allImages := false

	if all == "true" {
		allImages = true
	}

	err := docker.PullImage(reference, allImages, username, password)

	if err != nil {
		return fmt.Sprintf("Error Downloading: %s", reference), err
	}

	return fmt.Sprintf("Downloaded: %s", reference), err
}

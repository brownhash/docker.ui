package internals

import (
	"./docker"
	"./modals"
	"encoding/json"
	"fmt"
	"net/http"
)

func images() ([]modals.Image, error) {
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

func ImageHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	if request.Method == "GET" {
		imageList, err := images()

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			fmt.Println(err)
		}

		jsonResponse, _ := json.Marshal(imageList)

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonResponse)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

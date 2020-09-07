package internals

import (
	"./docker"
	"./errors"
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

	return imageList, err
}

func ImageHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	if request.Method == "GET" {
		imageList, err := images()

		if err != nil {
			errors.InternalServerError(w, err)
		}

		jsonResponse, _ := json.Marshal(imageList)

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonResponse)

		if err != nil {
			errors.InternalServerError(w, err)
		}

	} else {
		errors.MethodNotAllowed(w)
	}
}

func ImageDeletionHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	if request.Method == "POST" {
		var data map[string]interface{}
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&data)

		if err != nil {
			fmt.Println(err)
		}

		deletionData, err := ConvertToDeletionData(data)

		if err != nil {
			errors.InternalServerError(w, err)
		}

		for _, id := range deletionData.ImageIds {
			_, err := docker.DeleteImage(id, deletionData.Force, deletionData.PruneChildren)

			if err != nil {
				errors.InternalServerError(w, err)
			}
		}

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	} else {
		errors.MethodNotAllowed(w)
	}
}
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

func ImageDeletionHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	if request.Method == "POST" {
		var data map[string][]string
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&data)

		if err != nil {
			fmt.Println(err)
		}

		for _, id := range data["imageIds"] {
			_, err := docker.DeleteImage(id)

			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				_, err = w.Write([]byte(err.Error()))

				if err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}

				return
			}
		}

		_, err = fmt.Fprintf(w, "ok")

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
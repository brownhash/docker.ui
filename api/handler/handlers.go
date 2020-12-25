package handler

import (
	"encoding/json"
	"github.com/sharma1612harshit/docker.ui/api/docker"
	"log"
	"net/http"
)

// handler for images api
func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	images, err := docker.GetImages(r.Header.Get("all"), r.Header.Get("filters"))

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	} else {
		imageResponse, err := json.Marshal(images)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
		} else {
			w.Write([]byte(imageResponse))
		}
	}
}

// handler for containers api
func ContainersHandler(w http.ResponseWriter, r *http.Request) {
	containers, err := docker.GetContainers()

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	} else {
		containerResponse, err := json.Marshal(containers)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
		} else {
			w.Write([]byte(containerResponse))
		}
	}
}

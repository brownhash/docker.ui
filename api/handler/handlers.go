package handler

import (
	"encoding/json"
	"github.com/sharma1612harshit/docker.ui/api/docker"
	"log"
	"net/http"
)

// handler for images api
func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	imageResponse, err := json.Marshal(docker.GetImages())

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	} else {
		w.Write([]byte(imageResponse))
	}
}

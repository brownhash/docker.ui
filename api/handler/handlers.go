package handler

import (
	"encoding/json"
	"fmt"
	"github.com/sharma1612harshit/docker.ui/api/docker"
	"github.com/sharma1612harshit/gomuf/logs"
	"net/http"
)

// ImagesHandler - handler for images api
func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		images, err := docker.GetImages(r.Header.Get("all"), r.Header.Get("filters"))

		if err != nil {
			logs.Warn(err)

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("500 - Something bad happened!"))
		} else {
			imageResponse, err := json.Marshal(images)

			if err != nil {
				logs.Warn(err)

				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("500 - Something bad happened!"))
			} else {
				_, _ = w.Write(imageResponse)
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method not allowed!"))
	}
}

// PullImageHandler - handler for pull image api
func PullImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		status, err := docker.PullImage(r.Header.Get("all"), r.Header.Get("imageref"), r.Header.Get("username"), r.Header.Get("password"))

		if err != nil {
			logs.Warn(err)

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("%s - %s", status, err)))
		} else {
			_, _ = w.Write([]byte(status))
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method not allowed!"))
	}
}

// ContainersHandler - handler for containers api
func ContainersHandler(w http.ResponseWriter, r *http.Request) {
	containers, err := docker.GetContainers()

	if err != nil {
		logs.Warn(err)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("%s", err)))
	} else {
		containerResponse, err := json.Marshal(containers)

		if err != nil {
			logs.Warn(err)

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("%s", err)))
		} else {
			_, _ = w.Write(containerResponse)
		}
	}
}

// LaunchContainerHandler - handler for run container api
func LaunchContainerHandler(w http.ResponseWriter, r *http.Request) {
	launchResponse, err := docker.RunContainer(r.Header.Get("name"), r.Header.Get("image_name"), "", "", "", []string{}, []string{}, map[string]string{})

	if err != nil {
		logs.Warn(err)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("%s", err)))
	} else {
		response, err := json.Marshal(launchResponse)

		if err != nil {
			logs.Warn(err)

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("%s", err)))
		} else {
			_, _ = w.Write(response)
		}
	}
}

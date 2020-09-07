package internals

import (
	"./docker"
	"./errors"
	"./modals"
	"encoding/json"
	"net/http"
)

func containers() ([]modals.Container, error) {
	containers, err := docker.ListContainers()

	if err != nil {
		return nil, err
	}

	var containerList []modals.Container

	for _, container := range containers {
		tmp := modals.Container{
			Id:              container.ID,
			ImageId:         container.ImageID,
			Labels:          container.Labels,
			State:           container.State,
			Status:          container.Status,
			Mounts:          container.Mounts,
			Names:           container.Names,
			Ports:           container.Ports,
			HostConfig:       container.HostConfig,
			NetworkSettings: container.NetworkSettings,
		}
		containerList = append(containerList, tmp)
	}

	return containerList, err
}

func ContainerHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	if request.Method == "GET" {
		containerList, err := containers()

		if err != nil {
			errors.InternalServerError(w, err)
		}

		jsonResponse, _ := json.Marshal(containerList)

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonResponse)

		if err != nil {
			errors.InternalServerError(w, err)
		}

	} else {
		errors.MethodNotAllowed(w)
	}
}

package docker

import (
	"github.com/sharma1612harshit/docker.ui/pkg/docker"
	"log"
)

// return containers data as json map
func GetContainers() (map[string][]ContainerResponse, error) {
	containers, err := docker.GetContainers(false, false, true, false, "", "", 0)

	var response = map[string][]ContainerResponse{}

	if err != nil {
		log.Fatal(err)
		return response, err
	}

	var containerList []ContainerResponse

	for _, data := range containers {
		containerList = append(containerList, ContainerResponse{
			ID:      			data.ID,
			Names:   			data.Names,
			Created: 			data.Created,
			Image:   			data.Image,
			ImageId: 			data.ImageID,
			Labels:  			data.Labels,
			Command: 			data.Command,
			Mounts:  			data.Mounts,
			Ports:   			data.Ports,
			State:   			data.State,
			Status:  			data.Status,
			SizeRw:  			data.SizeRw,
			NetworkSettings: 	data.NetworkSettings,
		})
	}

	response["containers"] = containerList

	return response, err
}

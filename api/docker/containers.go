package docker

import (
	duiTypes "github.com/sharma1612harshit/docker.ui/api/types"
	"github.com/sharma1612harshit/docker.ui/pkg/docker"
	"log"
)

// return containers data as json map
func GetContainers() ([]duiTypes.ContainerResponse, error) {
	containers, err := docker.GetContainers(false, false, true, false, "", "", 0)

	var containerList = make([]duiTypes.ContainerResponse, 0)

	if err != nil {
		log.Print(err)
		return containerList, err
	}

	for _, data := range containers {
		containerList = append(containerList, duiTypes.ContainerResponse{
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

	return containerList, err
}

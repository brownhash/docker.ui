package docker

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	duiTypes "github.com/sharma1612harshit/docker.ui/api/types"
	"github.com/sharma1612harshit/docker.ui/pkg/docker"
	"github.com/sharma1612harshit/docker.ui/pkg/logger"
)

// return containers data as json map
func GetContainers() ([]duiTypes.ContainerResponse, error) {
	containers, err := docker.GetContainers(false, false, true, false, "", "", 0)

	var containerList = make([]duiTypes.ContainerResponse, 0)

	if err != nil {
		logger.Warn(err)
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

// run a container with specified configuration
func RunContainer(containerName, imageName, hostName, domainName, user string, cmd, entryPoint []string, labels map[string]string) (container.ContainerCreateCreatedBody, error) {
	containerConf := &container.Config{
		Hostname:        hostName,
		Domainname:      domainName,
		User:            user,
		Cmd:             cmd,
		Image:           imageName,
		Entrypoint:      entryPoint,
		Labels:          labels,
		//Env:             "",
		//Tty:             tty,
		//AttachStdin:     false,
		//AttachStdout:    false,
		//AttachStderr:    false,
		//ExposedPorts:    "",
		//NetworkDisabled: false,
		//MacAddress:      "",
		//OnBuild:         nil,
		//Volumes:         "",
		//WorkingDir:      "",
		//StopSignal:      "",
		//StopTimeout:     nil,
		//Shell:           nil,
		//OpenStdin:       false,
		//StdinOnce:       false,
		//ArgsEscaped:     false,
	}

	hostConf := &container.HostConfig{
		//Binds:           nil,
		//ContainerIDFile: "",
		//LogConfig:       container.LogConfig{},
		//NetworkMode:     "",
		//PortBindings:    nil,
		//RestartPolicy:   container.RestartPolicy{},
		//AutoRemove:      false,
		//VolumeDriver:    "",
		//VolumesFrom:     nil,
		//CapAdd:          nil,
		//CapDrop:         nil,
		//DNS:             nil,
		//DNSOptions:      nil,
		//DNSSearch:       nil,
		//ExtraHosts:      nil,
		//GroupAdd:        nil,
		//IpcMode:         "",
		//Cgroup:          "",
		//Links:           nil,
		//OomScoreAdj:     0,
		//PidMode:         "",
		//Privileged:      false,
		//PublishAllPorts: false,
		//ReadonlyRootfs:  false,
		//SecurityOpt:     nil,
		//StorageOpt:      nil,
		//Tmpfs:           nil,
		//UTSMode:         "",
		//UsernsMode:      "",
		//ShmSize:         0,
		//Sysctls:         nil,
		//Runtime:         "",
		//ConsoleSize:     [2]uint{},
		//Isolation:       "",
		//Resources:       container.Resources{},
		//Mounts:          nil,
		//Init:            nil,
		//InitPath:        "",
	}

	netConf := &network.NetworkingConfig{}

	resp, err := docker.LaunchContainer(containerConf, hostConf, netConf, containerName)

	return resp, err
}

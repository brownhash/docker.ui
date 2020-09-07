package modals

import (
	"github.com/docker/docker/api/types"
)

type Container struct {
	Id 					string								`json:"id"`
	ImageId 			string								`json:"image_id"`
	Labels 				map[string]string					`json:"labels"`
	State 				string								`json:"state"`
	Status 				string								`json:"status"`
	Mounts 				[]types.MountPoint					`json:"mounts"`
	Names 				[]string							`json:"names"`
	Ports 				[]types.Port						`json:"ports"`
	HostConfig 			interface{}							`json:"host_config"`
	NetworkSettings 	*types.SummaryNetworkSettings		`json:"network_settings"`
}

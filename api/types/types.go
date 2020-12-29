package types

import (
	"github.com/docker/docker/api/types"
)

// Structure to capture container response
type ContainerResponse struct {
	ID 					string 							`json:"id"`
	Names 				[]string 						`json:"names"`
	Created 			int64  							`json:"created"`
	Image 				string 							`json:"image"`
	ImageId 			string 							`json:"image_id"`
	Labels 				map[string]string 				`json:"labels"`
	Command 			string 							`json:"command"`
	Mounts 				[]types.MountPoint 				`json:"mounts"`
	Ports 				[]types.Port 					`json:"ports"`
	State 				string 							`json:"state"`
	Status 				string 							`json:"status"`
	SizeRw 				int64 							`json:"size_rw"`
	NetworkSettings 	*types.SummaryNetworkSettings 	`json:"network_settings"`
}

// Structure to capture image response
type ImageResponse struct {
	ID 			string 				`json:"id"`
	Created 	int64  				`json:"created"`
	Containers 	int64 				`json:"containers"`
	Labels 		map[string]string 	`json:"labels"`
	RepoDigest	[]string 			`json:"repo_digest"`
	RepoTags	[]string 			`json:"repo_tags"`
	Size 		int64				`json:"size"`
	ParentId 	string 				`json:"parent_id"`
}

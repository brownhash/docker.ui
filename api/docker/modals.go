package docker

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
package modals

type Image struct {
	Id 			string					`json:"id"`
	RepoTag 	[]string				`json:"repo_tag"`
	Size 		int64					`json:"size"`
	RepoDigest 	[]string				`json:"repo_digest"`
	CreatedAt 	int64					`json:"created_at"`
	Labels 		map[string]string		`json:"labels"`
	Containers 	int64					`json:"containers"`
}

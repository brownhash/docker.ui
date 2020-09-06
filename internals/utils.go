package internals

import (
	"./modals"
	"encoding/json"
)

func ConvertToDeletionData(requestData map[string]interface{}) (modals.ImageDeletionData, error) {
	result := modals.ImageDeletionData{
		Force:         false,
		PruneChildren: false,
		ImageIds:      nil,
	}
	var deletionChecks map[string]bool
	var imageIds []string

	checksJson, err := json.Marshal(requestData["checks"])

	if err != nil {
		return result, err
	}

	err = json.Unmarshal(checksJson, &deletionChecks)

	if err != nil {
		return result, err
	}

	imagesJson, err := json.Marshal(requestData["imageIds"])

	if err != nil {
		return result, err
	}

	err = json.Unmarshal(imagesJson, &imageIds)

	if err != nil {
		return result, err
	}

	result.Force = deletionChecks["force"]
	result.PruneChildren = deletionChecks["pruneChildren"]
	result.ImageIds = imageIds

	return result, err

}

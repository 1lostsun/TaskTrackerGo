package validation

import (
	"errors"
	"strings"
	"taskTrackerGo/internal/model"
)

func TaskGroupUpdatesBuilder(request model.TaskGroupRequest) (map[string]interface{}, error) {
	updates := map[string]interface{}{}

	if request.GroupLead != nil {
		if strings.TrimSpace(*request.GroupLead) == "" {
			return nil, errors.New("group lead cannot be empty")
		}
		updates["group_lead"] = *request.GroupLead
	}

	if request.Name != nil {
		if strings.TrimSpace(*request.Name) == "" {
			return nil, errors.New("name cannot be empty")
		}

		updates["name"] = *request.Name
	}

	return updates, nil
}

package validation

import (
	"errors"
	"fmt"
	"strings"
	"taskTrackerGo/internal/model"
	"time"
)

func IsValidState(status string) bool {
	switch status {
	case "pending":
		return true
	case "in_progress":
		return true
	case "done":
		return true
	case "overdue":
		return true
	default:
		return false
	}
}

func TaskUpdatesBuilder(request model.TaskRequest) (map[string]interface{}, error) {
	updates := map[string]interface{}{}

	if request.GroupID != nil {
		if *request.GroupID == 0 {
			return nil, errors.New("groupID must be greater than zero")
		}
		updates["group_id"] = *request.GroupID
	}
	if request.Name != nil {
		if strings.TrimSpace(*request.Name) == "" {
			return nil, errors.New("name cannot be empty")
		}
		updates["name"] = *request.Name
	}
	if request.Description != nil {
		updates["description"] = *request.Description
	}
	if request.TaskState != nil {
		if !IsValidState(*request.TaskState) {
			return nil, fmt.Errorf("invalid task state: %s", *request.TaskState)
		}
		updates["task_state"] = *request.TaskState
	}
	if request.Worker != nil {
		if strings.TrimSpace(*request.Worker) == "" {
			return nil, errors.New("worker cannot be empty")
		}
		updates["worker"] = *request.Worker
	}
	if request.Deadline != nil {
		if request.Deadline.Before(time.Now()) {
			return nil, errors.New("deadline cannot be in the past")
		}
		updates["deadline"] = *request.Deadline
	}

	if len(updates) == 0 {
		return nil, errors.New("updates must contain at least one item")
	}

	return updates, nil
}

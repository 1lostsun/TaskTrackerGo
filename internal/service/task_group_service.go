package service

import (
	"context"
	"taskTrackerGo/internal/model"
	"taskTrackerGo/internal/repository/postgres"
)

type TaskGroupService interface {
	CreateTaskGroup(ctx context.Context, taskGroup *model.TaskGroup) error
	GetTaskGroupByID(ctx context.Context, id uint64) (*model.TaskGroup, error)
	GetTaskGroupList(ctx context.Context) ([]*model.TaskGroup, error)
	UpdateTaskGroup(ctx context.Context, id uint64, updates map[string]interface{}) error
}

type taskGroupService struct {
	tgr postgres.TaskGroupRepository
}

func NewTaskGroupService(tgr postgres.TaskGroupRepository) TaskGroupService {
	return &taskGroupService{tgr}
}

func (tgs *taskGroupService) CreateTaskGroup(ctx context.Context, taskGroup *model.TaskGroup) error {
	err := tgs.tgr.CreateGroup(ctx, taskGroup)
	if err != nil {
		return err
	}

	return nil
}

func (tgs *taskGroupService) GetTaskGroupByID(ctx context.Context, id uint64) (*model.TaskGroup, error) {
	group, err := tgs.tgr.FindGroupByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (tgs *taskGroupService) GetTaskGroupList(ctx context.Context) ([]*model.TaskGroup, error) {
	groups, err := tgs.tgr.FindGroups(ctx)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (tgs *taskGroupService) UpdateTaskGroup(ctx context.Context, id uint64, updates map[string]interface{}) error {
	err := tgs.tgr.UpdateGroup(ctx, id, updates)
	if err != nil {
		return err
	}

	return nil
}

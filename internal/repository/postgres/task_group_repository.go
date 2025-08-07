package postgres

import (
	"context"
	"gorm.io/gorm"
	"taskTrackerGo/internal/model"
)

type TaskGroupRepository interface {
	CreateGroup(ctx context.Context, group *model.TaskGroup) error
	FindByID(ctx context.Context, id uint64) (*model.TaskGroup, error)
	FindAll(ctx context.Context) ([]*model.TaskGroup, error)
}

type taskGroupRepository struct {
	db *gorm.DB
}

func NewTaskGroupRepository(db *gorm.DB) TaskGroupRepository {
	return &taskGroupRepository{db}
}

func (r *taskGroupRepository) CreateGroup(ctx context.Context, group *model.TaskGroup) error {
	return r.db.WithContext(ctx).Create(group).Error
}

func (r *taskGroupRepository) FindByID(ctx context.Context, id uint64) (*model.TaskGroup, error) {
	var group model.TaskGroup
	err := r.db.WithContext(ctx).First(&group, id).Error
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *taskGroupRepository) FindAll(ctx context.Context) ([]*model.TaskGroup, error) {
	var taskGroups []*model.TaskGroup
	err := r.db.WithContext(ctx).Find(&taskGroups).Error
	if err != nil {
		return nil, err
	}

	return taskGroups, nil
}

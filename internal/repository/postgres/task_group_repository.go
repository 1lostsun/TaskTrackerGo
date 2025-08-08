package postgres

import (
	"context"
	"gorm.io/gorm"
	"taskTrackerGo/internal/model"
)

type TaskGroupRepository interface {
	CreateGroup(ctx context.Context, group *model.TaskGroup) error
	FindGroupByID(ctx context.Context, id uint64) (*model.TaskGroup, error)
	FindGroups(ctx context.Context) ([]*model.TaskGroup, error)
	UpdateGroup(ctx context.Context, id uint64, updates map[string]interface{}) error
	DeleteGroup(ctx context.Context, id uint64) error
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

func (r *taskGroupRepository) FindGroupByID(ctx context.Context, id uint64) (*model.TaskGroup, error) {
	var group model.TaskGroup
	err := r.db.WithContext(ctx).
		Preload("Tasks").
		First(&group, id).Error
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *taskGroupRepository) FindGroups(ctx context.Context) ([]*model.TaskGroup, error) {
	var taskGroups []*model.TaskGroup
	err := r.db.WithContext(ctx).
		Preload("Tasks").
		Order("id asc").
		Find(&taskGroups).Error
	if err != nil {
		return nil, err
	}

	return taskGroups, nil
}

func (r *taskGroupRepository) UpdateGroup(ctx context.Context, id uint64, updates map[string]interface{}) error {
	err := r.db.WithContext(ctx).
		Model(&model.TaskGroup{}).
		Where("id = ?", id).
		Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *taskGroupRepository) DeleteGroup(ctx context.Context, id uint64) error {
	err := r.db.WithContext(ctx).Delete(&model.TaskGroup{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

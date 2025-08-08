package postgres

import (
	"context"
	"gorm.io/gorm"
	"taskTrackerGo/internal/model"
	"time"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *model.Task) error
	FindTaskByID(ctx context.Context, id uint64) (*model.Task, error)
	FindTasksByGroupID(ctx context.Context, groupID uint64) ([]*model.Task, error)
	FindTasksByWorker(ctx context.Context, worker string) ([]*model.Task, error)
	FindOverdueTasksByGroupID(ctx context.Context, groupID uint64) ([]*model.Task, error)
	UpdateTask(ctx context.Context, id uint64, updates map[string]interface{}) error
	FindOverdueAndActiveTasks(ctx context.Context) ([]*model.Task, error)
	DeleteTaskByID(ctx context.Context, id uint64) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) CreateTask(ctx context.Context, task *model.Task) error {
	err := r.db.WithContext(ctx).Create(&task).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) FindTaskByID(ctx context.Context, id uint64) (*model.Task, error) {
	var task model.Task
	err := r.db.WithContext(ctx).First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepository) FindTasksByGroupID(ctx context.Context, groupID uint64) ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.db.WithContext(ctx).
		Where("group_id = ?", groupID).
		Order("id asc").
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) FindTasksByWorker(ctx context.Context, worker string) ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.db.WithContext(ctx).
		Where("worker = ?", worker).
		Order("created_at asc").
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) FindOverdueTasksByGroupID(ctx context.Context, groupID uint64) ([]*model.Task, error) {
	var overdueTasks []*model.Task
	err := r.db.WithContext(ctx).
		Where("group_id = ? AND deadline < ?", groupID, time.Now()).
		Order("created_at asc").
		Find(&overdueTasks).Error
	if err != nil {
		return nil, err
	}
	return overdueTasks, nil
}

func (r *taskRepository) FindOverdueAndActiveTasks(ctx context.Context) ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.db.WithContext(ctx).
		Order("id asc").
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, id uint64, updates map[string]interface{}) error {
	err := r.db.WithContext(ctx).
		Model(&model.Task{}).
		Where("id = ?", id).
		Updates(&updates).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) DeleteTaskByID(ctx context.Context, id uint64) error {
	err := r.db.WithContext(ctx).Delete(&model.Task{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

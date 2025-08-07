package postgres

import (
	"context"
	"gorm.io/gorm"
	"taskTrackerGo/internal/model"
	"time"
)

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
	FindByID(ctx context.Context, id uint64) (*model.Task, error)
	FindAllByGroupID(ctx context.Context, groupID uint64) ([]*model.Task, error)
	FindAllByWorker(ctx context.Context, worker string) ([]*model.Task, error)
	FindAllOverdue(ctx context.Context, groupID uint64) ([]*model.Task, error)
	Update(ctx context.Context, id uint64, updates map[string]interface{}) error
	DeleteByID(ctx context.Context, id uint64) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) Create(ctx context.Context, task *model.Task) error {
	err := r.db.WithContext(ctx).Create(&task).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) FindByID(ctx context.Context, id uint64) (*model.Task, error) {
	var task model.Task
	err := r.db.WithContext(ctx).First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepository) FindAllByGroupID(ctx context.Context, groupID uint64) ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.db.WithContext(ctx).Where("group_id = ?", groupID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) FindAllByWorker(ctx context.Context, worker string) ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.db.WithContext(ctx).Where("worker = ?", worker).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) FindAllOverdue(ctx context.Context, groupID uint64) ([]*model.Task, error) {
	var overdueTasks []*model.Task
	err := r.db.WithContext(ctx).
		Where("group_id = ? AND deadline < ?", groupID, time.Now()).
		Find(&overdueTasks).Error
	if err != nil {
		return nil, err
	}
	return overdueTasks, nil
}

func (r *taskRepository) Update(ctx context.Context, id uint64, updates map[string]interface{}) error {
	err := r.db.WithContext(ctx).
		Model(&model.Task{}).
		Where("id = ?", id).
		Updates(&updates).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) DeleteByID(ctx context.Context, id uint64) error {
	err := r.db.WithContext(ctx).Delete(&model.Task{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

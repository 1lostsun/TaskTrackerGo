package service

import (
	"context"
	"taskTrackerGo/internal/model"
	"taskTrackerGo/internal/repository/postgres"
)

type TaskService interface {
	CreateTask(ctx context.Context, task *model.Task) error
	GetTaskByID(ctx context.Context, id uint64) (*model.Task, error)
	GetAllTasksByGroupID(ctx context.Context, groupID uint64) ([]*model.Task, error)
	GetOverdueTasks(ctx context.Context, groupID uint64) ([]*model.Task, error)
	GetTasksByWorker(ctx context.Context, worker string) ([]*model.Task, error)
	UpdateTask(ctx context.Context, id uint64, updates map[string]interface{}) error
	DeleteTask(ctx context.Context, id uint64) error
}

type taskService struct {
	tr postgres.TaskRepository
}

func NewTaskService(tr postgres.TaskRepository) TaskService {
	return &taskService{tr}
}

func (s *taskService) CreateTask(ctx context.Context, task *model.Task) error {
	err := s.tr.Create(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) GetTaskByID(ctx context.Context, id uint64) (*model.Task, error) {
	task, err := s.tr.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) GetAllTasksByGroupID(ctx context.Context, groupID uint64) ([]*model.Task, error) {
	tasks, err := s.tr.FindAllByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *taskService) GetOverdueTasks(ctx context.Context, groupID uint64) ([]*model.Task, error) {
	overdueTasks, err := s.tr.FindAllOverdue(ctx, groupID)
	if err != nil {
		return nil, err
	}

	return overdueTasks, nil
}

func (s *taskService) GetTasksByWorker(ctx context.Context, worker string) ([]*model.Task, error) {
	workerTasks, err := s.tr.FindAllByWorker(ctx, worker)
	if err != nil {
		return nil, err
	}

	return workerTasks, nil
}

func (s *taskService) UpdateTask(ctx context.Context, id uint64, updates map[string]interface{}) error {
	err := s.tr.Update(ctx, id, updates)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) DeleteTask(ctx context.Context, id uint64) error {
	err := s.tr.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

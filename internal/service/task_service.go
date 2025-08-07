package service

import (
	"context"
	"taskTrackerGo/internal/model"
	"taskTrackerGo/internal/repository/postgres"
	"time"
)

type TaskService interface {
	CreateTask(ctx context.Context, task *model.Task) error
	GetTaskByID(ctx context.Context, id uint64) (*model.Task, error)
	GetTasksByGroupID(ctx context.Context, groupID uint64) ([]*model.Task, error)
	GetOverdueTasksByGroupID(ctx context.Context, groupID uint64) ([]*model.Task, error)
	GetTasksByWorker(ctx context.Context, worker string) ([]*model.Task, error)
	UpdateTask(ctx context.Context, id uint64, updates map[string]interface{}) error
	DeleteTask(ctx context.Context, id uint64) error
	EscalateOverdueTasks(ctx context.Context) error
}

type taskService struct {
	tr postgres.TaskRepository
}

func NewTaskService(tr postgres.TaskRepository) TaskService {
	return &taskService{tr}
}

func (s *taskService) CreateTask(ctx context.Context, task *model.Task) error {
	err := s.tr.CreateTask(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) GetTaskByID(ctx context.Context, id uint64) (*model.Task, error) {
	task, err := s.tr.FindTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) GetTasksByGroupID(ctx context.Context, groupID uint64) ([]*model.Task, error) {
	tasks, err := s.tr.FindTasksByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *taskService) GetOverdueTasksByGroupID(ctx context.Context, groupID uint64) ([]*model.Task, error) {
	overdueTasks, err := s.tr.FindOverdueTasksByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}

	return overdueTasks, nil
}

func (s *taskService) GetTasksByWorker(ctx context.Context, worker string) ([]*model.Task, error) {
	workerTasks, err := s.tr.FindTasksByWorker(ctx, worker)
	if err != nil {
		return nil, err
	}

	return workerTasks, nil
}

func (s *taskService) UpdateTask(ctx context.Context, id uint64, updates map[string]interface{}) error {
	err := s.tr.UpdateTask(ctx, id, updates)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) DeleteTask(ctx context.Context, id uint64) error {
	err := s.tr.DeleteTaskByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) EscalateOverdueTasks(ctx context.Context) error {
	updates := map[string]interface{}{}
	tasks, err := s.tr.FindOverdueAndActiveTasks(ctx)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if task.Deadline.Before(time.Now()) {
			if !(task.TaskState == "done") {
				updates["task_state"] = "overdue"
				err := s.tr.UpdateTask(ctx, task.ID, updates)
				if err != nil {
					return err
				}
			}
			continue
		}
	}

	return nil
}

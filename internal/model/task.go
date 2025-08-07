package model

import "time"

type Task struct {
	ID          uint64 `gorm:"primarykey"`
	GroupID     uint64
	Name        string
	Description string
	TaskState   string
	Worker      string
	CreatedAt   time.Time
	Deadline    time.Time
}

type TaskRequest struct {
	GroupID     uint64    `json:"group_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TaskState   string    `json:"task_state"`
	Worker      string    `json:"worker"`
	Deadline    time.Time `json:"deadline"`
}

type UpdateTaskRequest struct {
	GroupID     *uint64    `json:"group_id"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	TaskState   *string    `json:"task_state"`
	Worker      *string    `json:"worker"`
	Deadline    *time.Time `json:"deadline"`
}

type TaskResponse struct {
	GroupID     uint64    `json:"group_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TaskState   string    `json:"task_state"`
	Worker      string    `json:"worker"`
	CreatedAt   time.Time `json:"created_at"`
	Deadline    time.Time `json:"deadline"`
}

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"taskTrackerGo/internal/model"
	"taskTrackerGo/internal/service"
	"taskTrackerGo/internal/validation"
	"time"
)

type TaskHandler interface {
	CreateTaskHandler(c *gin.Context)
	GetTaskByIDHandler(c *gin.Context)
	GetTasksByGroupIDHandler(c *gin.Context)
	GetOverdueTasksByGroupIDHandler(c *gin.Context)
	GetTasksByWorkerHandler(c *gin.Context)
	UpdateTaskHandler(c *gin.Context)
	DeleteTaskHandler(c *gin.Context)
}

type taskHandler struct {
	ts service.TaskService
}

func NewTaskHandler(ts service.TaskService) TaskHandler {
	return &taskHandler{ts: ts}
}

func (h *taskHandler) CreateTaskHandler(c *gin.Context) {
	ctx := c.Request.Context()

	var input model.TaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := model.Task{
		GroupID:     *input.GroupID,
		Name:        *input.Name,
		Description: *input.Description,
		TaskState:   *input.TaskState,
		Worker:      *input.Worker,
		Deadline:    *input.Deadline,
		CreatedAt:   time.Now(),
	}

	err := h.ts.CreateTask(ctx, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "task created successfully"})
}

func (h *taskHandler) GetTaskByIDHandler(c *gin.Context) {
	ctx := c.Request.Context()
	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.ts.GetTaskByID(ctx, taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	taskResponse := model.TaskResponse{
		GroupID:     task.GroupID,
		Name:        task.Name,
		Description: task.Description,
		TaskState:   task.TaskState,
		Worker:      task.Worker,
		Deadline:    task.Deadline,
		CreatedAt:   task.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"task": taskResponse})
}

func (h *taskHandler) GetTasksByGroupIDHandler(c *gin.Context) {
	ctx := c.Request.Context()
	groupIDStr := c.Param("id")
	groupID, err := strconv.ParseUint(groupIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks, err := h.ts.GetTasksByGroupID(ctx, groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tasksResponse := make([]model.TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		tasksResponse = append(tasksResponse, model.TaskResponse{
			GroupID:     task.GroupID,
			Name:        task.Name,
			Description: task.Description,
			TaskState:   task.TaskState,
			Worker:      task.Worker,
			Deadline:    task.Deadline,
			CreatedAt:   task.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasksResponse})
}

func (h *taskHandler) GetOverdueTasksByGroupIDHandler(c *gin.Context) {
	ctx := c.Request.Context()
	groupIDStr := c.Param("id")
	groupID, err := strconv.ParseUint(groupIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks, err := h.ts.GetOverdueTasksByGroupID(ctx, groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tasksResponse := make([]model.TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		tasksResponse = append(tasksResponse, model.TaskResponse{
			GroupID:     task.GroupID,
			Name:        task.Name,
			Description: task.Description,
			TaskState:   task.TaskState,
			Worker:      task.Worker,
			Deadline:    task.Deadline,
			CreatedAt:   task.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasksResponse})
}

func (h *taskHandler) GetTasksByWorkerHandler(c *gin.Context) {
	ctx := c.Request.Context()
	workerStr := c.Param("worker")

	tasks, err := h.ts.GetTasksByWorker(ctx, workerStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tasksResponse := make([]model.TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		tasksResponse = append(tasksResponse, model.TaskResponse{
			GroupID:     task.GroupID,
			Name:        task.Name,
			Description: task.Description,
			TaskState:   task.TaskState,
			Worker:      task.Worker,
			Deadline:    task.Deadline,
			CreatedAt:   task.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasksResponse})
}

func (h *taskHandler) UpdateTaskHandler(c *gin.Context) {
	var input model.TaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates, err := validation.TaskUpdatesBuilder(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.ts.UpdateTask(ctx, taskID, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task updated successfully"})
}

func (h *taskHandler) DeleteTaskHandler(c *gin.Context) {
	ctx := c.Request.Context()
	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.ts.DeleteTask(ctx, taskID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}

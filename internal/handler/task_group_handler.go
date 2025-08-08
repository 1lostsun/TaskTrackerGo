package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"taskTrackerGo/internal/model"
	"taskTrackerGo/internal/service"
	"taskTrackerGo/internal/validation"
)

type TaskGroupHandler interface {
	CreateTaskGroupHandler(c *gin.Context)
	GetTaskGroupByIDHandler(c *gin.Context)
	GetTaskGroupListHandler(c *gin.Context)
	UpdateTaskGroupHandler(c *gin.Context)
}

type taskGroupHandler struct {
	tgs service.TaskGroupService
}

func NewTaskGroupHandler(tgs service.TaskGroupService) TaskGroupHandler {
	return &taskGroupHandler{tgs}
}

func (h *taskGroupHandler) CreateTaskGroupHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var taskGroupResponse model.TaskGroupRequest
	if err := c.ShouldBindJSON(&taskGroupResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskGroup := model.TaskGroup{
		Name:      *taskGroupResponse.Name,
		GroupLead: *taskGroupResponse.GroupLead,
		Tasks:     []model.Task{},
	}

	err := h.tgs.CreateTaskGroup(ctx, &taskGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "task group was successfully created"})
}

func (h *taskGroupHandler) GetTaskGroupByIDHandler(c *gin.Context) {
	ctx := c.Request.Context()
	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskGroup, err := h.tgs.GetTaskGroupByID(ctx, ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task group": taskGroup})
}

func (h *taskGroupHandler) GetTaskGroupListHandler(c *gin.Context) {
	ctx := c.Request.Context()

	groups, err := h.tgs.GetTaskGroupList(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task groups": groups})
}

func (h *taskGroupHandler) UpdateTaskGroupHandler(c *gin.Context) {
	ctx := c.Request.Context()
	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var taskGroupRequest model.TaskGroupRequest
	if err := c.ShouldBindJSON(&taskGroupRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates, err := validation.TaskGroupUpdatesBuilder(taskGroupRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.tgs.UpdateTaskGroup(ctx, ID, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task group was successfully updated"})
}

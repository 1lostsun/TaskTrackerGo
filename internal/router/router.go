package router

import (
	"github.com/gin-gonic/gin"
	"taskTrackerGo/internal/handler"
)

type Router interface {
	InitRoutes()
}

type router struct {
	gin *gin.Engine
	th  handler.TaskHandler
}

func NewRouter(gin *gin.Engine, th handler.TaskHandler) Router {
	return &router{gin: gin, th: th}
}

func (r *router) InitRoutes() {
	v1 := r.gin.Group("api/v1")
	{
		v1.POST("/tasks", r.th.CreateTaskHandler)
		v1.GET("/tasks/:id", r.th.GetTaskByIDHandler)
		v1.GET("/tasks/group/:id", r.th.GetAllTasksByGroupIDHandler)
		v1.GET("tasks/overdue/:id", r.th.GetOverdueTasksByGroupIDHandler)
		v1.GET("/tasks/worker/:worker", r.th.GetTasksByWorkerHandler)
		v1.PATCH("/tasks/:id", r.th.UpdateTaskHandler)
		v1.DELETE("tasks/:id", r.th.DeleteTaskHandler)
	}
}

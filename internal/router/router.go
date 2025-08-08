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
	tgh handler.TaskGroupHandler
}

func NewRouter(gin *gin.Engine, th handler.TaskHandler, tgh handler.TaskGroupHandler) Router {
	return &router{gin: gin, th: th, tgh: tgh}
}

func (r *router) InitRoutes() {
	v1 := r.gin.Group("api/v1")
	{
		v1.POST("/tasks", r.th.CreateTaskHandler)
		v1.GET("/tasks/:id", r.th.GetTaskByIDHandler)
		v1.GET("/tasks/group/:id", r.th.GetTasksByGroupIDHandler)
		v1.GET("tasks/overdue/:id", r.th.GetOverdueTasksByGroupIDHandler)
		v1.GET("/tasks/worker/:worker", r.th.GetTasksByWorkerHandler)
		v1.PATCH("/tasks/:id", r.th.UpdateTaskHandler)
		v1.DELETE("tasks/:id", r.th.DeleteTaskHandler)
	}

	v2 := r.gin.Group("api/v2")
	{
		v2.POST("/task_groups", r.tgh.CreateTaskGroupHandler)
		v2.GET("/task_groups/:id", r.tgh.GetTaskGroupByIDHandler)
		v2.GET("/task_groups", r.tgh.GetTaskGroupListHandler)
		v2.PATCH("task_groups/:id", r.tgh.UpdateTaskGroupHandler)
	}
}

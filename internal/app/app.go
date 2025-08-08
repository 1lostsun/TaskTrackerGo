package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"taskTrackerGo/internal/handler"
	"taskTrackerGo/internal/repository/postgres"
	"taskTrackerGo/internal/router"
	"taskTrackerGo/internal/scheduler"
	"taskTrackerGo/internal/service"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dotenvErr := godotenv.Load(".env")
	if dotenvErr != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := postgres.NewDSN()
	database, err := postgres.NewDBConnect(dsn)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	taskGroupRepo := postgres.NewTaskGroupRepository(database)
	taskRepo := postgres.NewTaskRepo(database)
	taskGroupService := service.NewTaskGroupService(taskGroupRepo)
	taskService := service.NewTaskService(taskRepo)
	scheduler.StartEscalationScheduler(ctx, taskService)
	taskGroupHandler := handler.NewTaskGroupHandler(taskGroupService)
	taskHandler := handler.NewTaskHandler(taskService)
	engine := gin.Default()
	r := router.NewRouter(engine, taskHandler, taskGroupHandler)
	r.InitRoutes()

	runErr := engine.Run(":8080")
	if runErr != nil {
		log.Fatal(runErr)
	}
}

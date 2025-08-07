package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"taskTrackerGo/internal/handler"
	"taskTrackerGo/internal/repository/postgres"
	"taskTrackerGo/internal/router"
	"taskTrackerGo/internal/service"
)

func Run() {
	dotenvErr := godotenv.Load(".env")
	if dotenvErr != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := postgres.NewDSN()
	database, err := postgres.NewDBConnect(dsn)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	taskRepo := postgres.NewTaskRepo(database)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)
	engine := gin.Default()
	r := router.NewRouter(engine, taskHandler)
	r.InitRoutes()

	runErr := engine.Run(":8080")
	if runErr != nil {
		log.Fatal(runErr)
	}
}

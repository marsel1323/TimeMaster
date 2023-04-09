package main

import (
	"TimeMaster/internal/api/handlers"
	"TimeMaster/internal/db"
	repository "TimeMaster/internal/repository/psql"
	"TimeMaster/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	dbConn, err := db.NewConnection()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer dbConn.Close()

	// Initialize handlers
	categoryHandler := handlers.NewCategoryHandler(service.NewCategoryService(repository.NewCategoryRepository(dbConn)))
	goalHandler := handlers.NewGoalHandler(service.NewGoalService(repository.NewGoalRepository(dbConn)))
	taskHandler := handlers.NewTaskHandler(service.NewTaskService(repository.NewTaskRepository(dbConn)))

	// Register routes
	e.GET("/categories", categoryHandler.GetAllCategories)
	e.GET("/categories/:id", categoryHandler.GetCategory)
	e.POST("/categories", categoryHandler.CreateCategory)

	e.GET("/goals", goalHandler.GetAllGoals)
	e.GET("/goals/:id", goalHandler.GetGoal)
	e.POST("/goals", goalHandler.CreateGoal)

	e.GET("/tasks", taskHandler.GetAllTasks)
	e.GET("/tasks/:id", taskHandler.GetTask)
	e.POST("/tasks", taskHandler.CreateTask)

	e.Logger.Fatal(e.Start(":8080"))
}

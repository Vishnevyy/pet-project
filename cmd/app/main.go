package main

import (
	"log"
	"pet-project/internal/database"
	"pet-project/internal/handlers"
	"pet-project/internal/taskService"
	"pet-project/internal/web/tasks"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()

	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)
	
	e := echo.New()
	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
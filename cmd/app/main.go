package main

import (
	"log"
	"pet-project/internal/database"
	"pet-project/internal/handlers"
	"pet-project/internal/taskService"
	"pet-project/internal/userService"
	"pet-project/internal/web/tasks"
	"pet-project/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()

	if err := database.DB.AutoMigrate(&taskService.Task{}, &userService.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	taskRepo := taskService.NewTaskRepository(database.DB)
	taskService := taskService.NewService(taskRepo)
	taskHandler := handlers.NewHandler(taskService)

	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}

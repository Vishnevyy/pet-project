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
	taskSvc := taskService.NewService(taskRepo)

	userRepo := userService.NewUserRepository(database.DB)
	userSvc := userService.NewUserService(userRepo)

	taskHandler := handlers.NewHandler(taskSvc, userSvc)
	userHandler := handlers.NewUserHandler(userSvc)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	strictTaskHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	log.Println("Starting server on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

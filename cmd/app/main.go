package main

import (
	"net/http"
	"pet-project/internal/database"
	"pet-project/internal/handlers"
	"pet-project/internal/taskService"

	"github.com/gorilla/mux"
)

func main() {
    database.InitDB()
 //   database.DB.AutoMigrate(&taskService.Task{})

    repo := taskService.NewTaskRepository(database.DB)
    service := taskService.NewService(repo)

    handler := handlers.NewHandler(service)

    router := mux.NewRouter()
    router.HandleFunc("/api/tasks", handler.GetTaskHandler).Methods("GET")
    router.HandleFunc("/api/tasks", handler.PostTaskHandler).Methods("POST")
    router.HandleFunc("/api/tasks/{id}", handler.PatchTaskHandler).Methods("PATCH")
    router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskHandler).Methods("DELETE")

    http.ListenAndServe(":8080", router)
}
package handlers

import (
	"context"

	"pet-project/internal/taskService"
	"pet-project/internal/web/tasks"
)

type Handler struct {
	Service taskService.TaskService
}

func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTask()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     new(int64),
			Task:   tsk.Title,
			IsDone: tsk.Completed,
		}
		*task.Id = int64(tsk.ID)
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := taskService.Task{
		Title:     taskRequest.Task,
		Completed: taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(&taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     new(int64),
		Task:   createdTask.Title,
		IsDone: createdTask.Completed,
	}
	*response.Id = int64(createdTask.ID)

	return response, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.Service.DeleteTaskByID(uint(request.Id))
	if err != nil {
		return tasks.DeleteTasksId404Response{}, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	updatedTask, err := h.Service.UpdateTaskByID(uint(request.Id), &taskService.Task{
		Title:     *request.Body.Task,
		Completed: *request.Body.IsDone,
	})
	if err != nil {
		return tasks.PatchTasksId404Response{}, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     new(int64),
		Task:   updatedTask.Title,
		IsDone: updatedTask.Completed,
	}
	*response.Id = int64(updatedTask.ID)

	return response, nil
}

func NewHandler(service taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

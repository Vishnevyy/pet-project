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
			Id:        new(int64),
			Title:     tsk.Title,
			Completed: tsk.Completed,
		}
		*task.Id = int64(tsk.ID)
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil {
		return tasks.PostTasks400Response{}, nil
	}

	if request.Body.Title == "" {
		return tasks.PostTasks400Response{}, nil
	}

	taskToCreate := taskService.Task{
		Title:     request.Body.Title,
		Completed: request.Body.Completed,
	}

	createdTask, err := h.Service.CreateTask(&taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:        new(int64),
		Title:     createdTask.Title,
		Completed: createdTask.Completed,
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
	if request.Body == nil {
		return tasks.PatchTasksId400Response{}, nil
	}

	if request.Body.Title == nil && request.Body.Completed == nil {
		return tasks.PatchTasksId400Response{}, nil
	}

	updatedTask, err := h.Service.UpdateTaskByID(uint(request.Id), &taskService.Task{
		Title:     *request.Body.Title,
		Completed: *request.Body.Completed,
	})
	if err != nil {
		return tasks.PatchTasksId404Response{}, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:        new(int64),
		Title:     updatedTask.Title,
		Completed: updatedTask.Completed,
	}
	*response.Id = int64(updatedTask.ID)

	return response, nil
}

func NewHandler(service taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

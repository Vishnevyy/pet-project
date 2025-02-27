package handlers

import (
	"context"

	"pet-project/internal/taskService"
	"pet-project/internal/userService"
	"pet-project/internal/web/tasks"
)

type Handler struct {
	TaskService taskService.TaskService
	UserService userService.UserService
}

func (h *Handler) GetUsersIdTasks(ctx context.Context, request tasks.GetUsersIdTasksRequestObject) (tasks.GetUsersIdTasksResponseObject, error) {
	userID := uint(request.Id)

	_, err := h.UserService.GetUserByID(userID)
	if err != nil {
		return tasks.GetUsersIdTasks404JSONResponse{Error: new(string)}, err
	}

	userTasks, err := h.TaskService.GetTasksForUser(userID)
	if err != nil {
		return tasks.GetUsersIdTasks404JSONResponse{Error: new(string)}, err
	}

	response := tasks.GetUsersIdTasks200JSONResponse{}
	for _, tsk := range userTasks {
		task := tasks.Task{
			Id:        new(int64),
			Title:     tsk.Title,
			Completed: tsk.Completed,
			UserId:    int64(tsk.UserID),
		}
		*task.Id = int64(tsk.ID)
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.TaskService.GetAllTask()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:        new(int64),
			Title:     tsk.Title,
			Completed: tsk.Completed,
			UserId:    int64(tsk.UserID),
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

	if request.Body.Title == "" || request.Body.UserId == 0 {
		return tasks.PostTasks400Response{}, nil
	}

	_, err := h.UserService.GetUserByID(uint(request.Body.UserId))
	if err != nil {
		return tasks.PostTasks400Response{}, nil
	}

	taskToCreate := taskService.Task{
		Title:     request.Body.Title,
		Completed: request.Body.Completed,
		UserID:    uint(request.Body.UserId),
	}

	createdTask, err := h.TaskService.CreateTask(&taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:        new(int64),
		Title:     createdTask.Title,
		Completed: createdTask.Completed,
		UserId:    int64(createdTask.UserID),
	}
	*response.Id = int64(createdTask.ID)

	return response, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.TaskService.DeleteTaskByID(uint(request.Id))
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

	existingTask, err := h.TaskService.GetTaskByID(uint(request.Id))
	if err != nil {
		return tasks.PatchTasksId404Response{}, err
	}

	if request.Body.Title != nil {
		existingTask.Title = *request.Body.Title
	}
	if request.Body.Completed != nil {
		existingTask.Completed = *request.Body.Completed
	}

	updatedTask, err := h.TaskService.UpdateTaskByID(uint(request.Id), existingTask)
	if err != nil {
		return tasks.PatchTasksId404Response{}, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:        new(int64),
		Title:     updatedTask.Title,
		Completed: updatedTask.Completed,
		UserId:    int64(updatedTask.UserID),
	}
	*response.Id = int64(updatedTask.ID)

	return response, nil
}

func NewHandler(taskService taskService.TaskService, userService userService.UserService) *Handler {
	return &Handler{
		TaskService: taskService,
		UserService: userService,
	}
}

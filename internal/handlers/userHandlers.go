package handlers

import (
	"context"
	"log"
	"pet-project/internal/userService"
	"pet-project/internal/web/users"

	"github.com/oapi-codegen/runtime/types"
)

type UserHandler struct {
	Service userService.UserService
}

func NewUserHandler(service userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		log.Printf("Failed to get users: %v", err)
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		user := users.UserResponse{
			Id:    int64(usr.ID),
			Email: types.Email(usr.Email),
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	if request.Body == nil {
		return users.PostUsers400Response{}, nil
	}

	userToCreate := userService.User{
		Email:    string(request.Body.Email),
		Password: request.Body.Password,
	}

	createdUser, err := h.Service.CreateUser(&userToCreate)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:    int64(createdUser.ID),
		Email: types.Email(createdUser.Email),
	}

	return response, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	if request.Body == nil {
		return users.PatchUsersId400Response{}, nil
	}

	updatedUser, err := h.Service.UpdateUserByID(uint(request.Id), &userService.User{
		Email:    string(*request.Body.Email),
		Password: *request.Body.Password,
	})
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:    int64(updatedUser.ID),
		Email: types.Email(updatedUser.Email),
	}

	return response, nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.Service.DeleteUserByID(uint(request.Id))
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}

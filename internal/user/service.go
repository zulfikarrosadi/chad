package user

import (
	"context"
	"net/http"
)

type UserCreateRequest struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profilePicture"`
}

type UserCreateResponse struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profilePicture"`
}

type UserUpdateRequest struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profilePicture"`
}

type UserUpdateResponse struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profilePicture"`
}

type Response struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   Error  `json:"error"`
}

type Error struct {
	Message string       `json:"message"`
	Details ErrorDetails `json:"details"`
}

type ErrorDetails struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type UserService interface {
	Get(ctx context.Context, userID int) *Response
	Create(ctx context.Context, user *UserCreateRequest) *Response
	Update(ctx context.Context, user *UserUpdateRequest) *Response
	Delete(ctx context.Context, userID int) *Response
}

type UserServiceImpl struct {
	Repository UserRepositoryImpl
}

func NewUserService(userRepository UserRepositoryImpl) UserService {
	return &UserServiceImpl{
		Repository: userRepository,
	}
}

func (us UserServiceImpl) Get(ctx context.Context, userID int) *Response {
	user, err := us.Repository.Get(ctx, userID)
	if err != nil {
		return &Response{
			Status: "error",
			Code:   http.StatusNotFound,
			Error: Error{
				Message: "user not found",
			},
		}
	}

	return &Response{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "user found",
		Data:    &user,
	}
}

func (us UserServiceImpl) Create(ctx context.Context, user *UserCreateRequest) *Response {
	userCreateResponse, err := us.Repository.Create(ctx, user)
	if err != nil {
		return &Response{
			Status: "error",
			Code:   http.StatusBadRequest,
			Error: Error{
				Message: "cannot create new user, please try again",
			},
		}
	}

	return &Response{
		Status:  "success",
		Code:    http.StatusCreated,
		Message: "user created",
		Data:    userCreateResponse,
	}
}

func (us UserServiceImpl) Update(ctx context.Context, user *UserUpdateRequest) *Response {
	userUpdateResponse, err := us.Repository.Update(ctx, user)
	if err != nil {
		return &Response{
			Status: "error",
			Code:   http.StatusBadRequest,
			Error: Error{
				Message: "cannot update user, please try again",
			},
		}
	}

	return &Response{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "user updated",
		Data:    userUpdateResponse,
	}
}

func (us UserServiceImpl) Delete(ctx context.Context, userID int) *Response {
	err := us.Repository.Delete(ctx, userID)
	if err != nil {
		return &Response{
			Status: "error",
			Code:   http.StatusBadRequest,
			Error: Error{
				Message: "cannot delete user, please try again",
			},
		}
	}
	return &Response{
		Code: http.StatusNoContent,
	}
}

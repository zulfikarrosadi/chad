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

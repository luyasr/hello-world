package user

import "context"

type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	GetUser(context.Context, int) (*User, error)
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

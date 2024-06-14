package user

import (
	"context"
	"fmt"
)

type UserService interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) error
	QueryUser(context.Context, *QueryRequest) (*User, error)
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Role:  RoleMember,
		Label: make(map[string]string),
	}
}

type CreateUserRequest struct {
	UserName string            `json:"username" gorm:"column:username"`
	Password string            `json:"password"`
	Role     Role              `json:"role"`
	Label    map[string]string `json:"label" gorm:"serializer:json"`
}

func (req *CreateUserRequest)ValidateAccount() error {
	if req.UserName == "" || req.Password == "" {
		return fmt.Errorf("username or password invalid")
	}
	return nil
}

func NewDeleteUserRequest(id int) *DeleteUserRequest {
	return &DeleteUserRequest{
		Id: id,
	}
}

type DeleteUserRequest struct {
	Id int `json:"id"`
}

func NewQueryRequestId(value string) *QueryRequest {
	return &QueryRequest{
		QueryBy: QueryById,
		QueryValue: value,
	}
}

func NewQueryRequestName(value string) *QueryRequest {
	return &QueryRequest{
		QueryBy: QueryByName,
		QueryValue: value,
	}
}

type QueryRequest struct {
	QueryBy QueryBy
	QueryValue string
}

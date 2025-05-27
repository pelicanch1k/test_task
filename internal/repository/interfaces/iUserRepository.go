package interfaces

import (
	"context"
	"test_task/internal/repository/gen"
)

type IUserRepository interface {
	GetUsers(ctx context.Context, filters *gen.GetUsersParams) (*[]gen.User, error)

	DeleteUser(ctx context.Context, id int32) error

	UpdateUser(ctx context.Context, userData *gen.UpdateUserParams) error

	CreateUser(ctx context.Context, user *gen.CreateUserParams) (*gen.User, error)
}
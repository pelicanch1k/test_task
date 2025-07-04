package interfaces

import (
	"context"
	"test_task/internal/model"
	"test_task/internal/repository/gen"
)

type IUserService interface {
	GetUsers(ctx context.Context, filters *model.GetUsersParams) ([]gen.User, error)

	GetUserByID(ctx context.Context,  id int32) (gen.User, error)

    DeleteUser(ctx context.Context, id int32) error

	UpdateUser(ctx context.Context, params *gen.UpdateUserParams) error

	CreateUser(ctx context.Context, user *model.CreateUserParams) (*gen.User, error)
}
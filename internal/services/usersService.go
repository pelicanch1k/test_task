package services

import (
	"context"
	"test_task/internal/repository"
	"test_task/internal/repository/gen"
	"test_task/internal/repository/interfaces"
)

type UsersService struct {
	usersRepository interfaces.IUserRepository
}


func NewUserService() *UsersService {
	return &UsersService{usersRepository: repository.NewUserRepository()}
}

func (u *UsersService) CreateUser(ctx context.Context, user *gen.CreateUserParams) (*gen.User, error) {
	return u.usersRepository.CreateUser(ctx, user)
}

func (u *UsersService) UpdateUser(ctx context.Context, userData *gen.UpdateUserParams) error {
	return u.usersRepository.UpdateUser(ctx, userData)
}

func (u *UsersService) DeleteUser(ctx context.Context, id int32) error {
	return u.usersRepository.DeleteUser(ctx, id)
}

func (u *UsersService) GetUsers(ctx context.Context, filters *gen.GetUsersParams) (*[]gen.User, error) {
	return u.usersRepository.GetUsers(ctx, filters)
}
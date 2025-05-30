package repository

import (
	"context"
	"fmt"
	"test_task/internal/repository/gen"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) GetUsers(ctx context.Context, filters *gen.GetUsersParams) ([]gen.User, error) {
	repo, release, err := GetQueriesFromPool(ctx, globalPool)
	if err != nil {
		return nil, repoError("GetUsers")
	}
	defer release()

	fmt.Println(filters)

	fmt.Printf("%+v\n", gen.GetUsersParams(*filters))
	users, err := repo.GetUsers(ctx, gen.GetUsersParams(*filters))

	fmt.Println("len: ", len(users))

	return users, err
}

func (u *UserRepository) GetUserByID(ctx context.Context,  id int32) (gen.User, error) {
	repo, release, err := GetQueriesFromPool(ctx, globalPool)
	if err != nil {
		return gen.User{}, repoError("GetUserByID")
	}
	defer release()

	return repo.GetUserByID(ctx, id)
}

func (u *UserRepository) DeleteUser(ctx context.Context, id int32) error{
	repo, release, err := GetQueriesFromPool(ctx, globalPool)
	if err != nil {
		return repoError("DeleteUser")
	}
	defer release()

	return repo.DeleteUser(ctx, id)
}

func (u *UserRepository) UpdateUser(ctx context.Context, userData *gen.UpdateUserParams) error {
	repo, release, err := GetQueriesFromPool(ctx, globalPool)
	if err != nil {
		return repoError("UpdateUser")
	}
	defer release()

	_, err = repo.UpdateUser(ctx, *userData)
	return err
}

func (u *UserRepository) CreateUser(ctx context.Context, user *gen.CreateUserParams) (*gen.User, error) {
	repo, release, err := GetQueriesFromPool(ctx, globalPool)
	if err != nil {
		return nil, repoError("CreateUser")
	}
	defer release()

	userrow, err := repo.CreateUser(ctx, *user)

	return &userrow, err

}

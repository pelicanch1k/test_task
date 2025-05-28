package services

import (
	"context"
	"test_task/internal/model"
	"test_task/internal/repository"
	"test_task/internal/repository/gen"
	"test_task/internal/repository/interfaces"

	// "github.com/gofiber/fiber/v2/log"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
)

type UsersService struct {
	usersRepository interfaces.IUserRepository
}

func NewUserService() *UsersService {
	return &UsersService{usersRepository: repository.NewUserRepository()}
}

func (u *UsersService) CreateUser(ctx context.Context, user *model.CreateUserParams) (*gen.User, error) {
	var patronymic, gender, nation pgtype.Text
	var age pgtype.Int4

	patronymic.Scan(user.Patronymic)
	age.Scan(GetRequest("https://api.agify.io/?name=Dmitriy")["age"])
	gender.Scan(GetRequest("https://api.genderize.io/?name=Dmitriy")["gender"])

	// Получаем доступ к вложенным структурам с проверками
	if country, ok := GetRequest("https://api.nationalize.io/?name=Dmitriy")["country"].([]interface{}); ok && len(country) > 0 {
		if firstCountry, ok := country[0].(map[string]interface{}); ok {
			if countryID, ok := firstCountry["country_id"].(string); ok {
				log.Printf("Country ID: %s\n", countryID)
				nation.Scan(countryID)
			} else {
				log.Println("Поле country_id не найдено или не является строкой")
			}
		} else {
			log.Println("Первый элемент country не является объектом")
		}
	} else {
		log.Println("Поле country не найдено или не является массивом")
	}

	return u.usersRepository.CreateUser(ctx, &gen.CreateUserParams{
		Name:        user.Name,
		Surname:     user.Surname,
		Patronymic:  patronymic,
		Age:         age,
		Gender:      gender,
		Nationality: nation,
	})
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

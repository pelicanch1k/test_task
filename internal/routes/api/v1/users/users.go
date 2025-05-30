package users

import (
	"strconv"
	"strings"
	"test_task/internal/consts"
	"test_task/internal/model"
	"test_task/internal/repository/gen"
	"test_task/internal/services"
	"test_task/internal/services/interfaces"
	"test_task/internal/validation"

	"github.com/gofiber/fiber/v2"
	bind "github.com/idan-fishman/fiber-bind"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog"
)

type Routes struct {
	userService interfaces.IUserService
}

func InitRoutes(router fiber.Router) {
	routes := Routes{
		userService: services.NewUserService(),
	}

	router.Post("/", bind.New(bind.Config{
		Validator: validation.GetValidator(),
		Source:    bind.JSON,
	}, &CreateUser{}),
		routes.createUserHandler,
	)

	router.Get("/filter", bind.New(bind.Config{
		Validator: validation.GetValidator(),
		Source:    bind.Params,
	}, &GetUsers{}), routes.getUsersHandler)

	// Исправлено: используем :id вместо :userId
	router.Delete("/:id", routes.deleteUserHandler)

	router.Put("/:id", bind.New(bind.Config{
		Validator: validation.GetValidator(),
		Source:    bind.JSON,
	}, &UpdateUser{}), routes.updateUserHandler)
}

// @Summary		Add a new user
// @Description	Creates a new user
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			request	body		CreateUser	true	"Add User Request"
// @Success		200		{string}	string			"OK"
// @Failure		400		{string}	string			"Bad Request - Invalid input"
// @Failure		500		{string}	string			"Internal Server Error - Database or server issue"
// @Router			/api/v1/users [post]
func (r *Routes) createUserHandler(c *fiber.Ctx) error {
	l := c.Locals(consts.RequestLogger).(*zerolog.Logger)
	body := c.Locals(bind.JSON).(*CreateUser)

	user := model.CreateUserParams{
		Name:       body.Name,
		Surname:    body.Surname,
		Patronymic: body.Patronymic,
	}

	l.Debug().Any("user", user).Send()

	userEntity, err := r.userService.CreateUser(c.UserContext(), &user)
	if err != nil {
		l.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}

	l.Info().Msg("user created")

	return c.JSON(&model.ModelResponseDto{Status: model.ResponseStatusOk, Model: userEntity})
}

// @Summary      Get users with filters and pagination
// @Description  Returns a list of users with optional filtering and pagination
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        name         query    string  false  "Filter by name (case insensitive)"
// @Param        surname      query    string  false  "Filter by surname (case insensitive)"
// @Param        age          query    int     false  "Minimum age filter (use with age_2 for range)"
// @Param        age_2        query    int     false  "Maximum age filter (use with age for range)"
// @Param        gender       query    string  false  "Filter by gender (male/female/other)"
// @Param        nationality  query    string  false  "Filter by nationality (2-letter country code)"
// @Param        limit        query    int     false  "Number of items per page (default: 10)"
// @Param        offset       query    int     false  "Offset for pagination (default: 1)"
// @Success		200		{string}	string			"OK"
// @Failure      400  {string}  string  "Bad Request - Invalid filter parameters"
// @Failure      500  {string}  string  "Internal Server Error - Database or server issue"
// @Router       /api/v1/users/filter [get]
func (r *Routes) getUsersHandler(c *fiber.Ctx) error {
	l := c.Locals(consts.RequestLogger).(*zerolog.Logger)

	// Получаем query-параметры
	queryParams := new(GetUsers)
	if err := c.QueryParser(queryParams); err != nil {
		l.Error().Err(err).Msg("failed to parse query parameters")
		return fiber.ErrBadRequest
	}

	users, err := r.userService.GetUsers(c.UserContext(), (*model.GetUsersParams)(queryParams))
	if err != nil {
		l.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}

	l.Info().Msg("users found")
	return c.JSON(&model.ModelResponseDto{Status: model.ResponseStatusOk, Model: users})
}


// @Summary		Delete user
// @Description	Delete user
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param        id   path  int  true  "User ID"
// @Success		200		{object}	model.ModelResponseDto
// @Failure		400		{string}	string	"Bad Request - Invalid input"
// @Failure		500		{string}	string	"Internal Server Error - Database or server issue"
// @Router       /api/v1/users/{id} [delete]
func (r *Routes) deleteUserHandler(c *fiber.Ctx) error {
	l := c.Locals(consts.RequestLogger).(*zerolog.Logger)

	// Получаем ID из пути
	idParam := c.Params("id")
	
	// Проверяем, что значение не содержит фигурных скобок (плейсхолдер)
	if strings.ContainsAny(idParam, "{}") {
		l.Error().Str("id", idParam).Msg("invalid user ID format - placeholder detected")
		return fiber.NewError(fiber.StatusBadRequest, "invalid user ID: placeholder not replaced")
	}

	// Конвертируем в число
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		l.Error().Str("id", idParam).Msg("invalid user ID format")
		return fiber.NewError(fiber.StatusBadRequest, "invalid user ID: must be positive integer")
	}

	if err := r.userService.DeleteUser(c.UserContext(), int32(id)); err != nil {
		l.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}

	l.Info().Msg("user deleted")
	return c.JSON(&model.ModelResponseDto{Status: model.ResponseStatusOk, Message: "ok"})
}

// ... (getUsersHandler остается без изменений) ...

// @Summary      Update user
// @Description  Updates user information
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Param        request body   UpdateUser  true  "User update data"
// @Success      200  {object}  model.ModelResponseDto
// @Failure      400  {string}  string  "Bad Request - Invalid input"
// @Failure      404  {string}  string  "Not Found - User not found"
// @Failure      500  {string}  string  "Internal Server Error - Database or server issue"
// @Router       /api/v1/users/{id} [put]
func (r *Routes) updateUserHandler(c *fiber.Ctx) error {
	l := c.Locals(consts.RequestLogger).(*zerolog.Logger)
	body := c.Locals(bind.JSON).(*UpdateUser)

	// Получаем ID из пути
	idParam := c.Params("id")
	
	// Проверяем, что значение не содержит фигурных скобок (плейсхолдер)
	if strings.ContainsAny(idParam, "{}") {
		l.Error().Str("id", idParam).Msg("invalid user ID format - placeholder detected")
		return fiber.NewError(fiber.StatusBadRequest, "invalid user ID: placeholder not replaced")
	}

	// Конвертируем в число
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		l.Error().Str("id", idParam).Msg("invalid user ID format")
		return fiber.NewError(fiber.StatusBadRequest, "invalid user ID: must be positive integer")
	}

	// Проверяем существование пользователя
	if _, err := r.userService.GetUserByID(c.UserContext(), int32(id)); err != nil {
		l.Error().Err(err).Msg("user not found")
		return fiber.ErrNotFound
	}

	// Подготавливаем параметры для обновления
	updateParams := gen.UpdateUserParams{
		ID:      int32(id),
	}

	if body.Name != nil {
		updateParams.Name = pgtype.Text{String: *body.Name, Valid: true}
	} else {
		updateParams.Name = pgtype.Text{Valid: false}
	}

	if body.Surname != nil {
		updateParams.Surname = pgtype.Text{String: *body.Surname, Valid: true}
	} else {
		updateParams.Surname = pgtype.Text{Valid: false}
	}

	if body.Patronymic != nil {
		updateParams.Patronymic = pgtype.Text{String: *body.Patronymic, Valid: true}
	} else {
		updateParams.Patronymic = pgtype.Text{Valid: false}
	}

	if body.Age != nil {
		updateParams.Age = pgtype.Int4{Int32: *body.Age, Valid: true}
	} else {
		updateParams.Age = pgtype.Int4{Valid: false}
	}

	if body.Gender != nil {
		updateParams.Gender = pgtype.Text{String: *body.Gender, Valid: true}
	} else {
		updateParams.Gender = pgtype.Text{Valid: false}
	}

	if body.Nationality != nil {
		updateParams.Nationality = pgtype.Text{String: *body.Nationality, Valid: true}
	} else {
		updateParams.Nationality = pgtype.Text{Valid: false}
	}

	// Обновляем данные
	err = r.userService.UpdateUser(c.UserContext(), &updateParams)
	if err != nil {
		l.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}

	user, err := r.userService.GetUserByID(c.UserContext(), int32(id))
	if err != nil {
		l.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}

	l.Info().Msg("user updated successfully")
	return c.JSON(&model.ModelResponseDto{
		Status: model.ResponseStatusOk,
		Model:  user,
	})
}
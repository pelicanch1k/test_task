package users

import (
	"test_task/internal/consts"
	"test_task/internal/model"
	"test_task/internal/services"
	"test_task/internal/services/interfaces"
	"test_task/internal/validation"

	"github.com/gofiber/fiber/v2"
	bind "github.com/idan-fishman/fiber-bind"
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

	router.Delete("/:userId", bind.New(bind.Config{
		Validator: validation.GetValidator(),
		Source:    bind.Params,
	}, &DeleteUser{}), routes.deleteUserHandler)

	router.Put("/:userId", bind.New(bind.Config{
		Validator: validation.GetValidator(),
		Source:    bind.Params,
	}, &DeleteUser{}), routes.deleteUserHandler)
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

// @Summary		Delete user
// @Description	Delete user
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param        id          query    int     true  "id юзера для удаления"
// @Success		200		{string}	string			"OK"
// @Failure		400		{string}	string			"Bad Request - Invalid input"
// @Failure		500		{string}	string			"Internal Server Error - Database or server issue"
// @Router /api/v1/users/{id} [delete]
func (r *Routes) deleteUserHandler(c *fiber.Ctx) error {
	l := c.Locals(consts.RequestLogger).(*zerolog.Logger)

	// Получаем query-параметры
	queryParams := new(DeleteUser)
	if err := c.QueryParser(queryParams); err != nil {
		l.Error().Err(err).Msg("failed to parse query parameters")
		return fiber.ErrBadRequest
	}

	if err := r.userService.DeleteUser(c.UserContext(), queryParams.Id); err != nil {
		l.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}

	l.Info().Msg("user deleted")
	return c.JSON(&model.ModelResponseDto{Status: model.ResponseStatusOk, Message: "ok"})
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

func (r *Routes) updateUsersHandler(c *fiber.Ctx) error {
	return nil
}
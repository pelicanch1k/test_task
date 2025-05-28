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

	router.Delete("/:userId", nil, routes.deleteUserHandler)

	router.Get("/", nil, routes.getUsersHandler)

	// router.Put("/:userId", )
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
		Name: body.Name,
		Surname: body.Surname,
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
// @Param id path string true "ID человека для удаления"
// @Success		200		{string}	string			"OK"
// @Failure		400		{string}	string			"Bad Request - Invalid input"
// @Failure		500		{string}	string			"Internal Server Error - Database or server issue"
// @Router /api/v1/users/{id} [delete]
func (r *Routes) deleteUserHandler(c *fiber.Ctx) error {
	l := c.Locals(consts.RequestLogger).(*zerolog.Logger)


	return nil
}

// @Summary		Get users
// @Description	Get users
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			request	body		CreateUser	true	"Get users"
// @Success		200		{string}	string			"OK"
// @Failure		400		{string}	string			"Bad Request - Invalid input"
// @Failure		500		{string}	string			"Internal Server Error - Database or server issue"
// @Router			/api/v1/users [get]
func (r *Routes) getUsersHandler(c *fiber.Ctx) error {

	return nil
}

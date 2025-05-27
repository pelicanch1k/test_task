package users

import (
	"test_task/internal/services/interfaces"

	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	userService interfaces.IUserService

	// codeStorage *utils.CodeStorage
}

func InitRoutes(router fiber.Router) {
	
}
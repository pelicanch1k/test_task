package v1

import (
	"test_task/internal/routes/api/v1/users"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(router fiber.Router) {
	users.InitRoutes(router.Group("/users"))
}

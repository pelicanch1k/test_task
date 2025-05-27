package routes

import (
	"test_task/internal/routes/api"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(router fiber.Router) {
	api.InitRoutes(router.Group("/api"))
}
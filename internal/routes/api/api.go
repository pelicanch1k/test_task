package api

import (
	_ "test_task/internal/routes/api/docs"
	v1 "test_task/internal/routes/api/v1"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title           lqd api
// @version         1.0
// @host            localhost:8080
// @description     This is a sample API.
// @securityDefinitions.apikey BearerAuth
// @in              header
// @name            Authorization
// @description     Enter your Bearer token here
func InitRoutes(router fiber.Router) {
	v1.InitRoutes(router.Group("/v1"))

	router.Get("/docs/*", fiberSwagger.WrapHandler)

}

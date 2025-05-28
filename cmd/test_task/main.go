package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"test_task/internal/config"
	"test_task/internal/consts"
	"test_task/internal/repository"
	"test_task/internal/routes"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339Nano,
	}).With().Caller().Logger()
}

func main() {

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := config.Load(ctxTimeout); err != nil {
		log.Fatal().Err(err)
	}

	pool, err := repository.GetConnectionPool(ctxTimeout,
		config.C.PostgresUsername, config.C.PostgresPassword,
		config.C.PostgresDb, "disable",
	)

	if err != nil {
		log.Fatal().Err(err)
	}

	repository.InitGlobalPgPool(pool)

	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(requestid.New(requestid.Config{
		Header:     fiber.HeaderXRequestID,
		Generator:  utils.UUID,
		ContextKey: consts.RequestId,
	}))

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger:      &log.Logger,
		WrapHeaders: true,
		Fields: []string{
			"ip", "latency", "status",
			"method", "url", "requestId",
			"reqHeaders", "ua", "body",
			"resBody", "queryParams", "error",
		},
		Next: func(c *fiber.Ctx) bool {
			return strings.HasPrefix(c.Path(), "/api/swagger/")
		},
	}))

	app.Use(func(c *fiber.Ctx) error {
		l := log.With().
			Str(consts.RequestId, c.Locals(consts.RequestId).(string)).
			Logger()

		c.Locals(consts.RequestLogger, &l)

		return c.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	routes.InitRoutes(app)

	addr := fmt.Sprintf(":%s", config.C.Port)
	log.Info().
		Str("port", config.C.Port).
		Msg("🚀 Server starting...")

	if err := app.Listen(addr); err != nil {
		log.Fatal().Err(err).Msg("Server failed")
	}}

package config

import (
	"context"
	"log"
	"test_task/internal/model"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

var C model.Config

func Load(ctx context.Context) error {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using system environment variables")
	}

	return envconfig.Process(ctx, &C)
}
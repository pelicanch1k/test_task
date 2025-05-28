package model

type Config struct {
	Port               string `env:"PORT"`
	PostgresUsername   string `env:"POSTGRES_USERNAME"`
	PostgresPassword   string `env:"POSTGRES_PASSWORD"`
	PostgresDb         string `env:"POSTGRES_DB"`
}
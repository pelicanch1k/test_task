package repository

import (
	"context"
	"errors"
	"fmt"
	"test_task/internal/repository/gen"

	"github.com/jackc/pgx/v5/pgxpool"
)

var globalPool *pgxpool.Pool

func InitGlobalPgPool(pool *pgxpool.Pool) {
	globalPool = pool
}

func GetConnectionPool(ctx context.Context, user, password, db, sslmode string) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, fmt.Sprintf(
		"postgresql://%s:%s@postgres:5432/%s?sslmode=%s",
		user, password, db, sslmode,
	))
}

func GetQueriesFromPool(ctx context.Context, pool *pgxpool.Pool) (*gen.Queries, func(), error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, nil, err
	}

	queries := gen.New(conn.Conn())
	release := func() { conn.Release() }

	return queries, release, nil
}

func repoError(method string) error {
	return errors.New(fmt.Sprintf("repository %s: repo, release, err := db.GetQueriesFromPool(ctx, globalPool)", method))
}

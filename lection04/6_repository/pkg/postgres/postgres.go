package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPool(dsn string) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

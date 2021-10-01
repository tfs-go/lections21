package repository

import (
	"context"
	"tfs-db/6_repository/domain"
	"tfs-db/6_repository/repository/queries"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repo struct {
	*queries.Queries
	pool *pgxpool.Pool
}

func NewRepository(pgxPool *pgxpool.Pool) Repository {
	return &repo{
		Queries: queries.New(pgxPool),
		pool:    pgxPool,
	}
}

type Repository interface {
	Candles(context.Context) ([]domain.Candle, error)
	CandlesByTicker(ctx context.Context, ticker string) ([]domain.Candle, error)
}

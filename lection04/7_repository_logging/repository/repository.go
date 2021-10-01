package repository

import (
	"context"
	"tfs-db/7_repository_logging/domain"
	"tfs-db/7_repository_logging/repository/queries"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type repo struct {
	*queries.Queries
	pool   *pgxpool.Pool
	logger logrus.FieldLogger
}

func NewRepository(pgxPool *pgxpool.Pool, logger logrus.FieldLogger) Repository {
	return &repo{
		Queries: queries.New(pgxPool),
		pool:    pgxPool,
		logger:  logger,
	}
}

type Repository interface {
	Candles(context.Context) ([]domain.Candle, error)
	CandlesByTicker(ctx context.Context, ticker string) ([]domain.Candle, error)
}

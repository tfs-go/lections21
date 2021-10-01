package services

import (
	"context"
	"tfs-db/9_clean_arch/internal/domain"

	"github.com/sirupsen/logrus"
)

type candlesRepo interface {
	Candles(context.Context) ([]domain.Candle, error)
}

type CandlesService struct {
	repo   candlesRepo
	logger logrus.FieldLogger
}

func NewCandlesService(logger logrus.FieldLogger, repo candlesRepo) *CandlesService {
	return &CandlesService{
		repo:   repo,
		logger: logger,
	}
}

func (c *CandlesService) Candles(ctx context.Context, filter domain.CandleFilter) ([]domain.Candle, error) {
	candles, err := c.repo.Candles(ctx)
	if err != nil {
		return nil, err
	}

	if len(candles) > filter.Limit && filter.Limit > 0 {
		candles = candles[:filter.Limit]
	}

	return candles, nil
}

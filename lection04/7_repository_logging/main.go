package main

import (
	"context"
	pkgpostgres "tfs-db/7_repository_logging/pkg/postgres"
	"tfs-db/7_repository_logging/repository"

	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)

	dsn := "postgres://user:passwd@localhost:5442/fintech" +
		"?sslmode=disable"

	pool, err := pkgpostgres.NewPool(dsn, logger)
	if err != nil {
		logger.Fatal(err)
	}
	defer pool.Close()

	repo := repository.NewRepository(pool, logger)

	if candles, err := repo.CandlesByTicker(context.Background(), "AAPL"); err != nil {
		logger.Error(err)
	} else {
		logger.Debugf("count: %d\n", len(candles))
		logger.Debugf("first candle: %+v", candles[0])
	}
}

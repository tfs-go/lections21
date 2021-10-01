package main

import (
	"context"
	"fmt"
	"log"
	"tfs-db/6_repository/repository"

	pkgpostgres "tfs-db/6_repository/pkg/postgres"
)

func main() {
	dsn := "postgres://user:passwd@localhost:5442/fintech" +
		"?sslmode=disable"

	pool, err := pkgpostgres.NewPool(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repo := repository.NewRepository(pool)

	candles, err := repo.CandlesByTicker(context.Background(), "AAPL")
	fmt.Printf("count: %d\n", len(candles))
	fmt.Printf("first candle: %+v", candles[0])
}
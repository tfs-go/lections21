package main

import (
	"net/http"
	"tfs-db/9_clean_arch/internal/handlers"
	"tfs-db/9_clean_arch/internal/repository"
	"tfs-db/9_clean_arch/internal/services"
	pkglog "tfs-db/9_clean_arch/pkg/log"
	pkgpostgres "tfs-db/9_clean_arch/pkg/postgres"

	"github.com/go-chi/chi/v5"
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

	candlesService := services.NewCandlesService(logger, repo)

	r := chi.NewRouter()
	r.Use(pkglog.NewStructuredLogger(logger))

	candlesHandler := handlers.NewCandles(logger, candlesService)
	r.Mount("/candles", candlesHandler.Routes())

	http.ListenAndServe(":3000", r)
}

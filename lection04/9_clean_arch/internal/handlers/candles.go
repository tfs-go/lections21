package handlers

import (
	"context"
	"net/http"
	"tfs-db/9_clean_arch/internal/domain"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

type CandlesService interface {
	Candles(ctx context.Context, filter domain.CandleFilter) ([]domain.Candle, error)
}

type Candles struct {
	logger  logrus.FieldLogger
	service CandlesService
}

func NewCandles(logger logrus.FieldLogger, service CandlesService) *Candles {
	return &Candles{
		logger:  logger,
		service: service,
	}
}

func (c *Candles) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", c.candles)
	})

	return r
}

func (c *Candles) candles(w http.ResponseWriter, r *http.Request) {
	var filter domain.CandleFilter
	filter.Bind(r)

	candles, err := c.service.Candles(r.Context(), filter)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.PlainText(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, candles)
}

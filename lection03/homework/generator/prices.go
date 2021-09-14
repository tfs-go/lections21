package generator

import (
	"context"
	"hw-async/domain"
	"math/rand"
	"time"
)

type Config struct {
	Factor  float64
	Delay   time.Duration
	Tickers []string
}

func NewPricesGenerator(cfg Config) *PricesGenerator {
	return &PricesGenerator{
		factor:  cfg.Factor,
		delay:   cfg.Delay,
		tickers: cfg.Tickers,
	}
}

type PricesGenerator struct {
	factor  float64
	delay   time.Duration
	tickers []string
}

func (g *PricesGenerator) Prices(ctx context.Context) <-chan domain.Price {
	prices := make(chan domain.Price)

	startTime := time.Now()
	go func() {
		defer close(prices)

		ticker := time.NewTicker(g.delay)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				currentTime := time.Now()
				delta := float64(currentTime.Unix() - startTime.Unix())
				ts := time.Unix(int64(float64(currentTime.Unix())+delta*g.factor), 0)

				for idx, ticker := range g.tickers {
					min := float64((idx + 1) * 100)
					max := min + 20
					prices <- domain.Price{
						Ticker: ticker,
						Value:  min + rand.Float64()*(max-min),
						TS:     ts,
					}
				}
			}
		}
	}()

	return prices
}

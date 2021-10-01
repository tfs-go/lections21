package domain

import (
	"net/http"
	"strconv"
	"time"
)

type CandlePeriod string

const (
	CandlePeriod1m  CandlePeriod = "1m"
	CandlePeriod2m  CandlePeriod = "2m"
	CandlePeriod10m CandlePeriod = "10m"
)

type Candle struct {
	Ticker string
	Period CandlePeriod
	Open   float64
	High   float64
	Low    float64
	Close  float64
	TS     time.Time
}

type CandleFilter struct {
	Limit int
}

func (f *CandleFilter) Bind(r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	var limit = 100
	if val, err := strconv.Atoi(limitStr); err == nil {
		limit = val
	}

	f.Limit = limit
}

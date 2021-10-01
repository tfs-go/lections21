package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
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

func main() {
	dsn := "postgres://user:passwd@localhost:5442/fintech" +
		"?sslmode=disable&fallback_application_name=fintech-app"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("can't connect to db: %s", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("can't ping db: %s", err)
	}

	const selectCandlesQuery = `SELECT instrument, period, ts, open, high, low,
		close FROM candles_1m`
	rows, err := db.QueryContext(context.Background(), selectCandlesQuery)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var candles []Candle
	for rows.Next() {
		var c Candle
		err = rows.Scan(&c.Ticker, &c.Period, &c.TS, &c.Open, &c.High, &c.Low, &c.Close)
		if err != nil {
			log.Fatal(err)
		}
		candles = append(candles, c)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

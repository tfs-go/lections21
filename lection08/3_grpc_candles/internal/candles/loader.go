package candles

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"tfs-grpc/3_grpc_candles/internal/domain"
	"time"
)

func LoadFromFile(prefix string, periods []domain.CandlePeriod) (map[domain.CandlePeriod]map[string][]domain.Candle, error) {
	var candlesStorage = map[domain.CandlePeriod]map[string][]domain.Candle{}
	for _, period := range periods {
		fileName := fmt.Sprintf("%s_%s.csv", prefix, period)
		rawCandles, err := getCandles(fileName)
		if err != nil {
			return nil, err
		}

		if candlesStorage[period], err = parseCandles(rawCandles, period); err != nil {
			return nil, err
		}
	}

	return candlesStorage, nil
}

func getCandles(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	candlesRaw, err := csv.NewReader(file).ReadAll()

	return candlesRaw, err
}

func parseCandles(candlesRaw [][]string, period domain.CandlePeriod) (map[string][]domain.Candle, error) {
	candles := map[string][]domain.Candle{}

	for _, candleRaw := range candlesRaw {
		candle, err := parseCandle(candleRaw)
		if err != nil {
			return candles, err
		}
		candle.Period = period

		candles[candle.Ticker] = append(candles[candle.Ticker], candle)
	}

	return candles, nil
}

func parseCandle(candleRaw []string) (candle domain.Candle, err error) {
	candle.Ticker = candleRaw[0]
	tsString := candleRaw[1]
	candle.TS, err = time.Parse(time.RFC3339, tsString)
	if err != nil {
		return
	}

	candle.Open, err = strconv.ParseFloat(candleRaw[2], 64)
	if err != nil {
		return
	}

	candle.High, err = strconv.ParseFloat(candleRaw[3], 64)
	if err != nil {
		return
	}

	candle.Low, err = strconv.ParseFloat(candleRaw[4], 64)
	if err != nil {
		return
	}

	candle.Close, err = strconv.ParseFloat(candleRaw[5], 64)
	if err != nil {
		return
	}

	return
}

package domain

import (
	"errors"
	"tfs-grpc/3_grpc_candles/internal/candlespb"
	"time"
)

type Candle struct {
	Ticker string
	Period CandlePeriod // Интервал
	Open   float64      // Цена открытия
	High   float64      // Максимальная цена
	Low    float64      // Минимальная цена
	Close  float64      // Цена закрытие
	TS     time.Time    // Время начала интервала
}

var ErrUnknownPeriod = errors.New("unknown period")

type CandlePeriod string

const (
	CandlePeriod1m  CandlePeriod = "1m"
	CandlePeriod2m               = "2m"
	CandlePeriod10m              = "10m"
	UnknownPeriod                = ""
)

func CandlePeriodFromProto(protoPeriod candlespb.Period) CandlePeriod {
	switch protoPeriod {
	case candlespb.Period_PERIOD_1M:
		return CandlePeriod1m
	case candlespb.Period_PERIOD_2M:
		return CandlePeriod2m
	case candlespb.Period_PERIOD_10M:
		return CandlePeriod10m
	default:
		return UnknownPeriod
	}
}

func CandlePeriodToProto(period CandlePeriod) (candlespb.Period, error) {
	switch period {
	case CandlePeriod1m:
		return candlespb.Period_PERIOD_1M, nil
	case CandlePeriod2m:
		return candlespb.Period_PERIOD_2M, nil
	case CandlePeriod10m:
		return candlespb.Period_PERIOD_10M, nil
	default:
		return candlespb.Period(0), ErrUnknownPeriod
	}
}

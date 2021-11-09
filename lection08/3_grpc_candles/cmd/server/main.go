package main

import (
	"log"
	"net"
	"tfs-grpc/3_grpc_candles/internal/candles"
	"tfs-grpc/3_grpc_candles/internal/candlespb"
	"tfs-grpc/3_grpc_candles/internal/domain"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	candlespb.UnsafeCandlesServiceServer
	candles map[domain.CandlePeriod]map[string][]domain.Candle
}

func (s *server) Candles(req *candlespb.CandleRequest, srv candlespb.CandlesService_CandlesServer) error {
	period := domain.CandlePeriodFromProto(req.Period)
	candlesStorage, ok := s.candles[period][req.Instrument]
	if !ok {
		return status.Error(codes.NotFound, "not found")
	}

	for _, candle := range candlesStorage {
		period, err := domain.CandlePeriodToProto(candle.Period)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if err := srv.Send(&candlespb.CandleResponse{
			Instrument: candle.Ticker,
			Ts:         timestamppb.New(candle.TS),
			Period:     period,
			Open:       candle.Open,
			High:       candle.High,
			Low:        candle.Low,
			Close:      candle.Close,
		}); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		time.Sleep(time.Second)
	}

	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":5005")
	if err != nil {
		log.Fatalf("can't listen on port: %v", err)
	}
	s := grpc.NewServer()

	candlesStorage, err := candles.LoadFromFile("candles",
		[]domain.CandlePeriod{domain.CandlePeriod1m, domain.CandlePeriod2m, domain.CandlePeriod10m})
	if err != nil {
		log.Fatalf("can't load candles: %v", err)
	}

	candlespb.RegisterCandlesServiceServer(s, &server{candles: candlesStorage})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("can't register service server: %v", err)
	}
}

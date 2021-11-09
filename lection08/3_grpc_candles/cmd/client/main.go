package main

import (
	"context"
	"fmt"
	"log"
	"tfs-grpc/3_grpc_candles/internal/candlespb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can't connect to server: %v", err)
	}
	client := candlespb.NewCandlesServiceClient(conn)
	fmt.Printf("created client: %v", client)

	client.Candles(context.Background(), &candlespb.CandleRequest{
		Instrument: "",
		Period:     0,
	})

	conn.Close()
}
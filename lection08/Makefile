gen-proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	$(PROTO_PATH)

candles:
	grpcurl -d '{"instrument": "AAPL", "period": 0}' \
	-proto 3_grpc_candles/internal/candlespb/CandlesService.proto -plaintext -v localhost:5005 candles.CandlesService/Candles
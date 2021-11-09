package main

import (
	"fmt"
	"log"
	"tfs-grpc/2_grpc_server_client/greetpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can't connect to server: %v", err)
	}
	client := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("created client: %v", client)
	conn.Close()
}

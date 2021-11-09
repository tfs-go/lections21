package main

import (
	"log"
	"net"
	"tfs-grpc/2_grpc_server_client/greetpb"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("can't listen on port: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, greetpb.UnimplementedGreetServiceServer{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("can't register service server: %v", err)
	}
}

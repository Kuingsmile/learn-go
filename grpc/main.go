package main

import (
	"grpcproj/proto"
	"grpcproj/server"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()
	proto.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

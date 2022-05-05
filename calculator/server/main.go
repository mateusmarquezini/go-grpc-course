package main

import (
	pb "github.com/mateusmarquezini/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen on: %v\n", err)
	}
	log.Printf("listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

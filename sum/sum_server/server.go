package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/carloslimasis/grpc-go-course/sum/sumpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, in *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	log.Printf("Received: %v", in.Sum)
	result := in.Sum.FirstNumber + in.Sum.SecondNumber
	return &sumpb.SumResponse{Result: result}, nil
}

func main() {

	fmt.Println("Start the sum gRPC Server...")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nirav24/grpc-go-service/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	message := fmt.Sprintf("Hello %s %s", request.Greeting.FirstName, request.Greeting.LastName)
	return &greetpb.GreetResponse{
		Result: message,
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

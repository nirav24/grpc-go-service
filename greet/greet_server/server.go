package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/nirav24/grpc-go-service/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	message := fmt.Sprintf("Hello %s %s", request.Greeting.FirstName, request.Greeting.LastName)
	return &greetpb.GreetResponse{
		Result: message,
	}, nil
}

func (s *server) GreetManyTimes(request *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := request.Greeting.GetFirstName()
	lastName := request.Greeting.GetLastName()
	for i := 1; i <= 10; i++ {
		message := fmt.Sprintf("Hello %s %s %d times", firstName, lastName, i)
		response := &greetpb.GreetManyTimesResponse{
			Result: message,
		}

		stream.Send(response)
		time.Sleep(1 * time.Second)
	}
	return nil
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

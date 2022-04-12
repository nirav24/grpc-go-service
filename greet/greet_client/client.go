package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nirav24/grpc-go-service/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect %v", conn)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	greetRequest := greetpb.GreetRequest{
		Greeting: &greetpb.Greetings{
			FirstName: "Nirav",
			LastName:  "Patel",
		},
	}

	response, err := c.Greet(context.Background(), &greetRequest)
	if err != nil {
		log.Printf("Error %v", err)
	}
	log.Printf("Response recieved %s ", response.Result)
}

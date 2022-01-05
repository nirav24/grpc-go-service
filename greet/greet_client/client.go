package main

import (
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

	fmt.Printf("Created client %f", c)
}

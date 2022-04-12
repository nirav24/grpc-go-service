package main

import (
	"context"
	"fmt"
	"io"
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
	log.Printf("Response received %s ", response.Result)

	greetManyTimesRequest := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greetings{
			FirstName: "Nirav",
			LastName:  "Patel",
		},
	}
	responseStream, err := c.GreetManyTimes(context.Background(), greetManyTimesRequest)
	if err != nil {
		log.Printf("Error recieving greet stream %v", err)
	}

	for {
		response, err := responseStream.Recv()
		if err == io.EOF {
			log.Println("greet stream is finished")
			break
		}
		if err != nil {
			log.Printf("Error recieving message from greet stream %v\n", err)
		}
		log.Printf("Response received %s\n", response.Result)
	}

}

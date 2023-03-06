package main

import (
	"context"
	"log"

	pb "github.com/berkayakcay/toy-go-grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("goGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Berkay",
	})

	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}

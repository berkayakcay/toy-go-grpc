package main

import (
	"context"
	"io"
	"log"

	pb "github.com/berkayakcay/toy-go-grpc/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Berkay",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v\n", err)

		}

		log.Printf("Greeting: %s\n", msg)
	}
}

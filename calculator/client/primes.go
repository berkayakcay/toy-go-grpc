package main

import (
	"context"
	"io"
	"log"

	pb "github.com/berkayakcay/toy-go-grpc/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Println("doPrime was invoked")

	req := &pb.PrimeRequest{
		Number: 1234567890,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v\n", err)

		}

		log.Printf("Primes: %d\n", res.Result)
	}
}

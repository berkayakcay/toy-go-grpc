package main

import (
	"io"
	"log"

	pb "github.com/berkayakcay/toy-go-grpc/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone function was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"
		stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatal("Error while sending data to client: %v\n", err)
		}
	}

}

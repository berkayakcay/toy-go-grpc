package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/berkayakcay/toy-go-grpc/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("error while reading cliend stream: %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}

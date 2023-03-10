package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/berkayakcay/toy-go-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {

	log.Printf("Sqrt was invoked with: %v\n", in)

	num := in.Number

	if num < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", num),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(num)),
	}, nil
}

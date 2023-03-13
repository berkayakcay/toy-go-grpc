package main

import (
	pb "github.com/berkayakcay/toy-go-grpc/blog/proto"
)

type Server struct {
	pb.BlogServiceServer
}

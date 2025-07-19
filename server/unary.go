package main

import (
	"context"
	pb "github.com/praneeth116/go-grpc/proto"
)

func (s *helloServer)SayHello(context.Context, *pb.NoParam)(*pb.HelloResponse, error){
	return &pb.HelloResponse{
		Message: "Hello Beyotch",
	}, nil
}
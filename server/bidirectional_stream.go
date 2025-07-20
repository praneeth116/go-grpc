package main

import (
	"io"
	"log"
	pb "github.com/praneeth116/go-grpc/proto"
)

func (s *helloServer)SayBiDirectionalStreaming(stream pb.GreetService_SayBiDirectionalStreamingServer)(error){
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil{
			return err
		}
		log.Printf("Got request with name %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello "+ req.Name,
		}
		if err := stream.Send(res);err != nil{
			return err
		}
	}
}
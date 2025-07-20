package main

import (
	"context"
	"log"
	"time"
	pb "github.com/praneeth116/go-grpc/proto"
)

func callSayHelloBiDirectionalStreaming(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Bidirectional streaming started!")
	stream, err := client.SayBiDirectionalStreaming(context.Background())
	if err != nil{
		log.Fatalf("couldn't send names %v", err)
	}
	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil{
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the request with name: %s", name)
		time.Sleep(2 * time.Second)

		res, err := stream.Recv()
		if err != nil{
			log.Fatalf("Error while receiving %v", err)
		}
		log.Printf("%v", res.Message)
	}
}
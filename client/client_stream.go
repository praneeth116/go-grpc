package main

import (
	"context"
	"log"
	"time"

	pb "github.com/praneeth116/go-grpc/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Client streaming started!")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil{
		log.Fatalf("couldn't send names: %v", err)
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
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished!")
	if err != nil{
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("%v", res.Messages)
}
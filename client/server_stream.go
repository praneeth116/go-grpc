package main

import (
	"context"
	"io"
	"log"

	pb "github.com/praneeth116/go-grpc/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList){
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil{
		log.Fatalf("could not send names: %v", err)
	}
	
	for {
		message, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Printf("error while streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}
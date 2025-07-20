package main

import (
	"context"
	"io"
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
	
	// Create a channel to signal when the receiving part is done.
	waitc := make(chan struct{})

	// Start a new goroutine to receive messages from the server concurrently.
	go func() {
		for{
			message, err := stream.Recv()
			if err == io.EOF{
				// If the server has finished sending, break the loop.
				break
			}
			if err != nil{
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message.Message)
		}
		// When the loop is done, close the channel to signal completion.
		close(waitc)
	}()

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req);err != nil{
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}

	// After sending all names, close the sending side of the stream.
	stream.CloseSend()

	// Wait here until the `waitc` channel is closed by the receiving goroutine.
	<-waitc

	log.Printf("Bidirectional streaming finished!")
}
package main

import (
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/praneeth116/go-grpc/proto"
)

const (
	port = ":8080"
)

func main(){
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	
	callSayHello(client)

}
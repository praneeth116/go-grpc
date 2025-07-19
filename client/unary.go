package main

import (
	"context"
	"log"
	pb "github.com/praneeth116/go-grpc/proto"
)

func callSayHello(client pb.GreetServiceClient){
	ctx := context.Background()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil{
		log.Fatalf("Failed to call the Unary api %e", err)
	}
	log.Printf("%s", res.Message)
}
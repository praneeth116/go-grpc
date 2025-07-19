package main

import (
	"log"
	"net"
	pb "github.com/praneeth116/go-grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct{
	pb.GreetServiceServer
}

func main(){
	lis, err := net.Listen("tcp",port) // Start listening for connections on our port.
	if err != nil{
		log.Fatalf("Failed to start the server %v",err)
	}
	grpcServer := grpc.NewServer() // Create a new gRPC server.

	// Register our service implementation with the gRPC server.
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{}) 
	log.Printf("server started at %v", lis.Addr())

	// Start the server and begin accepting client connections. This runs forever.
	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Failed to start %v", err)
	}
}
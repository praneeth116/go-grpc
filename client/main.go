package main

import (
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/praneeth116/go-grpc/proto"
)

const (
	// The address of the server we want to connect to.
	port = ":8080"
)

func main(){
	// Create a connection to the gRPC server.
	// We use `insecure.NewCredentials()` because we are not using SSL/TLS for this example.
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	// `defer` schedules this to run when the `main` function finishes, ensuring the connection is closed.
	defer conn.Close()

	// Create a new client "stub" using the connection. This is our "remote control" for the server.
	client := pb.NewGreetServiceClient(conn)
	
	// Call a function that uses our new client to make an RPC call.
	// callSayHello(client) //Unary RPC

	names := &pb.NameList{
		Names: []string{"Praneeth" ,"Pavan", "Akshay", "Aravind"},
	}

	callSayHelloServerStreaming(client, names) //Server streaming RPC

}
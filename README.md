# Go gRPC Showcase

This project is a practical demonstration of the four types of communication available in gRPC, implemented in Go. It serves as a clear, functional example for understanding Unary, Server Streaming, Client Streaming, and Bidirectional Streaming RPCs.

---
## 🚀 Features

This project implements a simple `GreetService` with four distinct RPC methods:

* **Unary RPC (`SayHello`)**: A classic request-response model where the client sends a single request and gets a single response.
* **Server Streaming RPC (`SayHelloServerStreaming`)**: The client sends a request with a list of names, and the server streams back a greeting for each name.
* **Client Streaming RPC (`SayHelloClientStreaming`)**: The client streams a sequence of names to the server, and the server responds with a single message summarizing the greetings.
* **Bidirectional Streaming RPC (`SayBiDirectionalStreaming`)**: The client and server can send messages to each other independently in a two-way communication stream.

---
## 📋 Prerequisites

Before you begin, ensure you have the following installed on your system:
* [Go](https://go.dev/doc/install) (version 1.18 or newer recommended)
* [Protocol Buffers Compiler (protoc)](https://grpc.io/docs/protoc-installation/)

---
## ⚙️ Setup & Installation

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/praneeth116/go-grpc.git
    cd go-grpc
    ```

2.  **Install Go dependencies for gRPC:**
    ```sh
    go get -u google.golang.org/grpc
    go get -u google.golang.org/protobuf/cmd/protoc-gen-go
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
    ```

3.  **Generate gRPC code from the `.proto` file:**
    This command compiles the `greet.proto` file and generates the necessary Go files (`greet.pb.go` and `greet_grpc.pb.go`) inside the `proto/` directory.
    ```sh
    protoc --go_out=. --go-grpc_out=. proto/greet.proto
    ```

4.  **Tidy Go modules:**
    ```sh
    go mod tidy
    ```

---
## ▶️ How to Run

You will need two separate terminal windows to run the server and the client.

### 1. Run the gRPC Server
In your first terminal, navigate to the server directory and run the `main.go` file.

```sh
cd server
go run .
```
You should see a confirmation message that the server has started:
```
INFO: server started at [::]:8080
```

### 2. Run the gRPC Client
To run a specific RPC, you need to edit the `client/main.go` file.

1.  **Open `client/main.go` in your code editor.**
2.  **Uncomment the line** for the RPC call you want to test and make sure the others are commented out.
3.  **Save the file.**
4.  **In your second terminal, run the client:**
    ```sh
    cd client
    go run .
    ```
You will see log messages in both the client and server terminals corresponding to the RPC call you enabled. To test another RPC type, simply edit and save the `client/main.go` file again and re-run the command.
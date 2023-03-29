# golang-grpc

* Prerequisite 
  - Protoc compiler
    - for Ubuntu:
      ```bash
      sudo apt-get install protobuf-compiler
      ```
    - for macOS:
      ```bash
      brew install protoc
      ```



1. Create one proto file name with `greeeing.proto` or can copy from [Here](https://github.com/DekivadiyaKishan/golang-grpc/blob/main/proto/greeting.proto)

```proto

syntax = "proto3";

option go_package = "/pb";

service GreetingService {
    rpc Greeting(GreetingServiceRequest) returns (GreetingServiceReply) {}
}

message GreetingServiceRequest {
    string name = 1;
}

message GreetingServiceReply {
    string message = 2;
}

```

2. Now We need to convert this proto file into go file

```bash
protoc <proto-file-path> --go_out=<output-file-path> --go-grpc_out=<output-file-path>
```

```bash
protoc *.proto --go_out=./ --go-grpc_out=./
```
  - after execute above command we can see two files are generated in `pb` directory
  - In `greeting_grpc.pb.go` one interface is auto generated 

```go
type GreetingServiceServer interface {
	Greeting(context.Context, *GreetingServiceRequest) (*GreetingServiceReply, error)
}
```

3. Now we need to inherite and implement above interface [Here](https://github.com/DekivadiyaKishan/golang-grpc/blob/main/server/server.go)

```go
package main

import (
	"context"
	"fmt"
	"grpc-golang/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.GreetingServiceServer
}

func (s *server) Greeting(ctx context.Context, req *pb.GreetingServiceRequest) (*pb.GreetingServiceReply, error) {
	return &pb.GreetingServiceReply{
		Message: fmt.Sprintf("Hello, %s", req.Name),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```
  - we implemented interface method with `server` struct
```go
 func (s *server) Greeting(ctx context.Context, req *pb.GreetingServiceRequest) (*pb.GreetingServiceReply, error) {
	return &pb.GreetingServiceReply{
		Message: fmt.Sprintf("Hello, %s", req.Name),
	}, nil
}
```
  - and register `server` struct with greeting server
```go
pb.RegisterGreetingServiceServer(s, &server{})
```
  - now we can start our server file with `go run server.go`

4. After server start we need to write client for call greeting method or copy from [Here](https://github.com/DekivadiyaKishan/golang-grpc/blob/main/client/client.go)

```go 
package main

import (
	"context"
	"fmt"
	"grpc-golang/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewGreetingServiceClient(cc)
	request := &pb.GreetingServiceRequest{Name: "Gophers"}

	resp, err := client.Greeting(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive response => %s ", resp.Message)
}
```
 - Here we use same auto generated file for create request and call greeting method
```go
client := pb.NewGreetingServiceClient(cc)
request := &pb.GreetingServiceRequest{Name: "Gophers"}

resp, err := client.Greeting(context.Background(), request)
if err != nil {
  log.Fatal(err)
}
```

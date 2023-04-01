package main

import (
	"context"
	"fmt"
	"go-rpc-practice/cmd/demogrpc/greeter/proto/greeter"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Hello struct {
}

func (h Hello) Hello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	fmt.Println(req)
	return &greeter.HelloResponse{Message: "Hello," + req.Name}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServer(grpcServer, new(Hello))

	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatalln("Listen:", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln("Serve:", err)
	}
}

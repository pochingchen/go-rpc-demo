package main

import (
	"context"
	"fmt"
	"go-rpc-practice/cmd/demogrpc/greeter/proto/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// credentials.NewClientTLSFromFile
	// grpc.WithTransportCredentials
	grpcClient, err := grpc.Dial("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Dial:", err)
	}
	client := greeter.NewGreeterClient(grpcClient)
	req := &greeter.HelloRequest{Name: "Kevin"}
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Response:", res)
}

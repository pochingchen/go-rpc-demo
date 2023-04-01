package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}
	defer conn.Close()

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("hello.Hello", "cdd", &reply)
	if err != nil {
		fmt.Println("Call:", err)
		return
	}
	fmt.Println("reply:", reply)

	var reply2 int
	err = client.Call("world.Increment", 100, &reply2)
	if err != nil {
		fmt.Println("Call:", err)
		return
	}
	fmt.Println("reply:", reply2)
}

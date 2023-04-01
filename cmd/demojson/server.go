package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Hello struct{}

func (h Hello) Hello(req string, res *string) error {
	fmt.Println("Hello:", req)
	*res = "hello " + req
	return nil
}

type World struct{}

func (w World) Increment(req int, res *int) error {
	fmt.Println("Increment:", req)
	*res = req + 1
	return nil
}

func main() {
	err := rpc.RegisterName("hello", new(Hello))
	if err != nil {
		fmt.Println("register name:", err)
		return
	}
	err = rpc.RegisterName("world", new(World))
	if err != nil {
		fmt.Println("register name:", err)
		return
	}
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Println("listen:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept:", err)
			return
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

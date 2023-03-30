package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Hello struct{}

func (h Hello) Hello(req string, res *string) error {
	*res = "hello " + req
	return nil
}

func (h Hello) What(req int, res *int) error {
	return nil
}

type World struct{}

func (w World) Increment(req int, res *int) error {
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
	listener, err := net.Listen("tcp", "127.0.0.1:9000")
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
		rpc.ServeConn(conn)
	}
}

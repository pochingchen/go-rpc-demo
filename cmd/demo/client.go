package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}
	defer conn.Close()

	var reply string
	err = conn.Call("hello.Hello", "cdd", &reply)
	if err != nil {
		fmt.Println("Call:", err)
		return
	}

	fmt.Println("reply:", reply)

	var reply2 int
	err = conn.Call("world.Increment", 100, &reply2)
	if err != nil {
		fmt.Println("Call:", err)
		return
	}

	fmt.Println("reply:", reply2)

	var reply3 int
	err = conn.Call("hello.What", 100, &reply2)
	if err != nil {
		fmt.Println("Call:", err)
		return
	}

	fmt.Println("reply:", reply3)
}

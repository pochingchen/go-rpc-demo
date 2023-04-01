package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Goods struct {
}

type AddGoodsRequest struct {
	ID      int
	Title   string
	Price   float32
	Content string
}

type AddGoodsResponse struct {
	Success bool
	Message string
}

type GetGoodsRequest struct {
	ID int
}

type GetGoodsResponse struct {
	ID      int
	Title   string
	Price   float32
	Content string
}

func (g Goods) AddGoods(req AddGoodsRequest, res *AddGoodsResponse) error {
	fmt.Println(req)
	*res = AddGoodsResponse{
		Success: true,
		Message: "add goods succeed",
	}
	return nil
}

func (g Goods) GetGoods(req GetGoodsRequest, res *GetGoodsResponse) error {
	fmt.Println(req)
	*res = GetGoodsResponse{
		ID:      1,
		Title:   "Xiaomi10",
		Price:   4999.0,
		Content: "Xiaomi10-5G",
	}
	return nil
}

func main() {
	err := rpc.RegisterName("goods", new(Goods))
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
	fmt.Println("Goods server start...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}

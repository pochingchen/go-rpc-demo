package main

import (
	"fmt"
	"net/rpc"
)

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

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Goods demo start...")

	var addGoodsReply AddGoodsResponse
	err = conn.Call("goods.AddGoods", AddGoodsRequest{
		ID:      1,
		Title:   "Xiaomi10",
		Price:   4999.0,
		Content: "Xiaomi10-5G",
	}, &addGoodsReply)
	if err != nil {
		fmt.Println("Call:", err)
		return
	}
	fmt.Println("add goods reply:", addGoodsReply)

	var getGoodsReply GetGoodsResponse
	err = conn.Call("goods.GetGoods", GetGoodsRequest{
		ID: 1,
	}, &getGoodsReply)
	if err != nil {
		fmt.Println("Call:", err)
		return
	}
	fmt.Println("get goods reply:", getGoodsReply)
}

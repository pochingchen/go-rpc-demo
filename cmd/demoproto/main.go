package main

import (
	"fmt"
	"go-rpc-practice/cmd/demoproto/proto/userinfo"
	"google.golang.org/protobuf/proto"
)

func main() {
	u := userinfo.Userinfo{
		Username: "cdd",
		Age:      20,
		Hobby:    []string{"CS-GO", "LOL"},
	}
	fmt.Println(u.GetUsername(), u.GetAge(), u.GetType(), u.GetHobby())

	data, err := proto.Marshal(&u)
	if err != nil {
		fmt.Println("proto marshal:", err)
		return
	}
	fmt.Println(data)

	var u2 userinfo.Userinfo
	err = proto.Unmarshal(data, &u2)
	if err != nil {
		fmt.Println("proto unmarshal:", err)
		return
	}
	fmt.Println(u2.Username, u2.Age, u2.Hobby)

}

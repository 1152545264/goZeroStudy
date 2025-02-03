package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type (
	GetUser struct {
		Id string
	}

	GetResp struct {
		Id    string
		Name  string
		Phone string
	}
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("建立连接失败", err)
	}
	defer client.Close()

	for i := 1; i <= 3; i++ {
		var (
			req  = GetUser{Id: fmt.Sprintf("%d", i)}
			resp GetResp
		)
		err = client.Call("UserServer.GetUser", req, &resp)
		if err != nil {
			log.Printf("第%d求失败, err:%v\n", i, err)
			continue
		}
		log.Println(resp)
	}
}

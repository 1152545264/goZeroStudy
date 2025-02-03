package main

import (
	"context"
	"fmt"
	"goZeroStudy/2testRpc/2-6grpc/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// grpc.Dial 已经废弃
	cli, err := grpc.NewClient("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("连接失败")
	}
	defer cli.Close()

	c := user.NewUserClient(cli)
	for i := 1; i <= 5; i++ {
		res, err := c.GetUser(context.Background(), &user.GetUserReq{Id: fmt.Sprintf("%d", i)})
		if err != nil {
			log.Printf("请求失败, err:%v\n", err)
			continue
		}
		log.Println(res)
	}
}

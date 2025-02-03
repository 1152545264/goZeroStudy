package main

import (
	"context"
	"errors"
	user "goZeroStudy/2testRpc/2-6grpc/proto/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserServer struct {
}

func (u *UserServer) GetUser(ctx context.Context, req *user.GetUserReq) (*user.GetUserResp, error) {
	if u, ok := users[req.Id]; ok {
		resp := &user.GetUserResp{
			Id: u.Id, Name: u.Name, Phone: u.Phone,
		}
		return resp, nil
	}
	return nil, errors.New("没有找到该用户")
}

func main() {
	lis, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("监听请求失败, err:%v\n", err)
	}
	s := grpc.NewServer()
	user.RegisterUserServer(s, new(UserServer))

	log.Println("服务已经启动....")
	s.Serve(lis)
}

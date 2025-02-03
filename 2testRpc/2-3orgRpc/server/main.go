package main

import (
	"errors"
	"log"
	"net"
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

type UserServer struct {
}

func (u *UserServer) GetUser(req GetUser, resp *GetResp) error {
	if u, ok := users[req.Id]; ok {
		*resp = GetResp{
			u.Id, u.Name, u.Phone,
		}
		return nil
	}
	return errors.New("没有找到该用户")
}

func main() {
	userServer := new(UserServer)

	rpc.Register(userServer)

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("监听失败", err)
	}
	log.Println("服务启动成功")

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("接受客户端连接失败", err)
			continue
		}

		go rpc.ServeConn(conn) // 异步处理客户端
	}
}

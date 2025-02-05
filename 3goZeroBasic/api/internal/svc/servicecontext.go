package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goZeroStudy/3goZeroBasic/api/internal/config"
	"goZeroStudy/3goZeroBasic/user/userclient"
)

type ServiceContext struct {
	Config config.Config

	UserClient userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserClient: userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
	}
}

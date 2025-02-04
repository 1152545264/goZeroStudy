package logic

import (
	"context"

	"goZeroStudy/3goZeroBasic/user/internal/svc"
	"goZeroStudy/3goZeroBasic/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 定义rpc方法
func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	// todo: add your logic here and delete this line

	if u, ok := users[in.Id]; ok {
		resp := &user.GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}
		return resp, nil
	}
	return &user.GetUserResp{}, nil
}

type User struct {
	Id    string
	Name  string
	Phone string
}

var users = map[string]*User{
	"1": {
		"1", "木兮", "123456",
	},
	"2": {
		"2", "兮木", "654321",
	},
}

package logic

import (
	"context"
	"goZeroStudy/3goZeroBasic/user/userclient"

	"goZeroStudy/3goZeroBasic/api/internal/svc"
	"goZeroStudy/3goZeroBasic/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserReq) (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line
	//resp = &types.UserResp{Id: "666", Name: "test", Phone: "testPhone"}

	getUserResp, err := l.svcCtx.UserClient.GetUser(l.ctx, &userclient.GetUserReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	res := &types.UserResp{
		Id:    getUserResp.Id,
		Name:  getUserResp.Name,
		Phone: getUserResp.Phone,
	}
	return res, nil
}

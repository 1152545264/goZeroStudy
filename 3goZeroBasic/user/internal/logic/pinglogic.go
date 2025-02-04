package logic

import (
	"context"

	"goZeroStudy/3goZeroBasic/user/internal/svc"
	"goZeroStudy/3goZeroBasic/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *user.GetUserReq) (*user.GetUserResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserResp{}, nil
}

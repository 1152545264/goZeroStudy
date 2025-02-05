package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"goZeroStudy/3goZeroBasic/user/internal/config"
	"goZeroStudy/3goZeroBasic/user/models"
)

type ServiceContext struct {
	Config    config.Config
	UserModel models.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: models.NewUserModel(sqlConn, c.Cache),
	}
}

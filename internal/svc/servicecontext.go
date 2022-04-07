package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gozero-test/internal/config"
	"gozero-test/model"
)

type ServiceContext struct {
	Config     config.Config
	ATestModel model.AtestModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ATestModel: model.NewAtestModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}

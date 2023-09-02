package svc

import (
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-app/rpc/auth/internal/config"
	"zero-app/rpc/auth/model"
)

type ServiceContext struct {
	Config config.Config
	Model  model.TestUsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewTestUsersModel(sqlx.NewSqlConn("postgres", c.DataSource)),
	}
}

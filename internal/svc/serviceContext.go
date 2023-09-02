package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"zero-app/internal/config"
	"zero-app/internal/middleware"
)

type ServiceContext struct {
	Config                 config.Config
	UserIdentityMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                 c,
		UserIdentityMiddleware: middleware.NewUserIdentityMiddleware().Handle,
	}
}

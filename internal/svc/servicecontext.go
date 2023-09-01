package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"zero-app/internal/config"
	"zero-app/internal/middleware"
	"zero-app/rpc/auth/auth"
)

type ServiceContext struct {
	Config                 config.Config
	UserIdentityMiddleware rest.Middleware
	AuthService            auth.AuthClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                 c,
		UserIdentityMiddleware: middleware.NewUserIdentityMiddleware().Handle,
	}
}

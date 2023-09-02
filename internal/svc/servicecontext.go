package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-app/internal/config"
	"zero-app/internal/middleware"
	"zero-app/rpc/auth/authClient"
)

type ServiceContext struct {
	Config                 config.Config
	UserIdentityMiddleware rest.Middleware
	Auth                   authClient.Auth
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                 c,
		UserIdentityMiddleware: middleware.NewUserIdentityMiddleware().Handle,
		Auth:                   authClient.NewAuth(zrpc.MustNewClient(c.Auth)),
	}
}

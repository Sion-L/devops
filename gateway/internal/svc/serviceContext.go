package svc

import (
	"github.com/Sion-L/devops/gateway/internal/config"
	"github.com/Sion-L/devops/gateway/internal/middleware"
	"github.com/Sion-L/devops/user/userClient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	User                 userClient.User
	AuthorizeMiddleware  rest.Middleware
	RefreshJwtMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		User:                 userClient.NewUser(zrpc.MustNewClient(c.User)),
		AuthorizeMiddleware:  middleware.NewAuthorizeMiddleware(c.Authorize.DataSource).Handle,
		RefreshJwtMiddleware: middleware.NewRefreshJwtMiddleware(c.Auth.AccessSecret, c.Auth.AccessExpire, c.Auth.TokenDisableDuration).Handle,
	}
}

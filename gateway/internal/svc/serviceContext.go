package svc

import (
	"github.com/Sion-L/devops/user/userClient"
	"github.com/Sion-L/gateway/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	User   userClient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userClient.NewUser(zrpc.MustNewClient(c.User)),
	}
}

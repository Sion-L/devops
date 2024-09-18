package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	User zrpc.RpcClientConf
	Auth struct {
		AccessSecret         string
		AccessExpire         int64
		TokenDisableDuration int64
	}
	Authorize struct {
		DataSource string
	}
	WebSocket struct {
		ReadBufferSize  int
		WriteBufferSize int
	}
}

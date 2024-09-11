package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Auth struct {
	AccessSecret string
	AccessExpire int64
}

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	Cache      cache.CacheConf
}

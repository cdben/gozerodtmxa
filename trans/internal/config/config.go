package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Mysql struct {
	Driver   string
	Host     string
	Port     int64
	User     string
	Password string
	Db       string
}

type Config struct {
	zrpc.RpcServerConf

	Mysql Mysql

	CacheRedis cache.CacheConf
}

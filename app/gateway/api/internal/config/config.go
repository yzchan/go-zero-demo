package config

import (
	"demo/lib"
	"github.com/zeromicro/go-queue/dq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DqConf        dq.DqConf
	RedisConf     *redis.RedisConf
	ConsulConf    *lib.ConsulConf
	ElasticSearch *lib.ElasticSearch

	Nsqlookupd struct {
		Hosts []string
	}
	Nsqd struct {
		Hosts   []string
		Topic   string
		Channel string
	}
}

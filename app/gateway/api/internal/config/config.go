package config

import (
	"demo/lib"
	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
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

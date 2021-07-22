package config

import (
	"demo/app/common"
	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DqConf        dq.DqConf
	RedisConf     *redis.RedisConf
	ConsulConf    *common.ConsulConf
	ElasticSearch *common.ElasticSearch

	Nsqlookupd struct {
		Hosts []string
	}
	Nsqd struct {
		Hosts   []string
		Topic   string
		Channel string
	}
}

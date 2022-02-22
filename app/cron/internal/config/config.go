package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	Debug     bool
	RedisConf redis.RedisConf
}

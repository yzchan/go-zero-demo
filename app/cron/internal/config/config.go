package config

import (
	"github.com/tal-tech/go-zero/core/stores/redis"
)

type Config struct {
	Debug     bool
	RedisConf redis.RedisConf
}

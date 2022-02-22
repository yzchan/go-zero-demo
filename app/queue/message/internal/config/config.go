package config

import (
	"github.com/zeromicro/go-queue/dq"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DqConf dq.DqConf
}

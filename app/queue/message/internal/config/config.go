package config

import (
	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DqConf dq.DqConf
}

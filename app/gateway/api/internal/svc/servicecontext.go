package svc

import (
	"demo/app/gateway/api/internal/config"
	"github.com/tal-tech/go-queue/dq"
)

type ServiceContext struct {
	Config   config.Config
	Producer dq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Producer: dq.NewProducer(c.DqConf.Beanstalks),
	}
}

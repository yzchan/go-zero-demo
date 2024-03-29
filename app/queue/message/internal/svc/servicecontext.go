package svc

import (
	"demo/app/queue/message/internal/config"
	"github.com/zeromicro/go-queue/dq"
)

type ServiceContext struct {
	Config     config.Config
	DqConsumer dq.Consumer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		DqConsumer: dq.NewConsumer(c.DqConf),
	}
}

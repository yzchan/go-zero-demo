package svc

import (
	"demo/app/services/iploc/internal/config"
	iplocate "github.com/yzchan/ip-locate"
)

type ServiceContext struct {
	Config config.Config
	IpLoc  *iplocate.QQWryParser
}

func NewServiceContext(c config.Config) *ServiceContext {
	parser, err := iplocate.NewQQWryParser(c.QQWryPath)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		IpLoc:  parser,
	}
}

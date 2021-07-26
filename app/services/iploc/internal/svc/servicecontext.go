package svc

import (
	"demo/app/services/iploc/internal/config"
	iploc "github.com/yzchan/iploc"
)

type ServiceContext struct {
	Config config.Config
	IpLoc  *iploc.QQWryParser
}

func NewServiceContext(c config.Config) *ServiceContext {
	parser, err := iploc.NewQQWryParser(c.QQWryPath)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		IpLoc:  parser,
	}
}

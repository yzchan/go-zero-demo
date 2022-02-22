package svc

import (
	"demo/app/gateway/tool/internal/config"
	"demo/app/gateway/tool/internal/middleware"
	"demo/app/services/iploc/iplocclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
	IpLoc   iplocclient.IpLoc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
		IpLoc:   iplocclient.NewIpLoc(zrpc.MustNewClient(c.IpLoc)),
	}
}

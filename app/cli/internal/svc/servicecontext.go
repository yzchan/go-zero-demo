package svc

import (
	"demo/app/cli/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
	}
}

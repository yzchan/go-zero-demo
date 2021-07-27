package svc

import (
	"demo/app/cron/internal/config"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
	//DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	//db, err := c.MysqlConf.MustGorm()
	//if err != nil {
	//	panic(err)
	//}

	return &ServiceContext{
		Config: c,
		Redis:  c.RedisConf.NewRedis(),
		//DB:     db,
	}
}

package svc

import (
	"demo/app/gateway/api/internal/config"
	"demo/app/gateway/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"

	"github.com/olivere/elastic/v7"
	"github.com/zeromicro/go-queue/dq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config        config.Config
	Example       rest.Middleware
	Producer      dq.Producer
	redisConf     *redis.RedisConf
	Redis         *redis.Redis
	ElasticSearch *elastic.Client
}

func NewServiceContext(c config.Config) *ServiceContext {

	// ElasticSearch
	auth := elastic.SetBasicAuth(c.ElasticSearch.Username, c.ElasticSearch.Password)
	esClient, err := elastic.NewClient(auth, elastic.SetSniff(false), elastic.SetURL(c.ElasticSearch.Hosts...))
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:        c,
		Example:       middleware.NewExampleMiddleware().Handle,
		Producer:      dq.NewProducer(c.DqConf.Beanstalks),
		Redis:         c.RedisConf.NewRedis(),
		ElasticSearch: esClient,
	}
}

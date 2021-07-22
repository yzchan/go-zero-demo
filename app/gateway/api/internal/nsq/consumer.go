package nsq

import (
	"demo/app/gateway/api/internal/nsq/handlers"
	"demo/app/gateway/api/internal/svc"
	"github.com/nsqio/go-nsq"
	"time"
)

func InitConsumer(ctx *svc.ServiceContext) (err error) {
	conf := nsq.NewConfig()
	conf.LookupdPollInterval = time.Second * 5
	conf.Snappy = true
	conf.MaxInFlight = 10

	var consumer *nsq.Consumer
	if consumer, err = nsq.NewConsumer(ctx.Config.Nsqd.Topic, ctx.Config.Nsqd.Channel, conf); err != nil {
		return
	}
	consumer.SetLoggerLevel(nsq.LogLevelError)

	handler := handlers.NewFetchHandler()
	consumer.AddConcurrentHandlers(handler, 10)
	if err = consumer.ConnectToNSQLookupds(ctx.Config.Nsqlookupd.Hosts); err != nil {
		return
	}
	return
}

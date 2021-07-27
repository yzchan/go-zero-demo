package main

import (
	"flag"
	"fmt"

	"demo/app/cron/internal/config"
	"demo/app/cron/internal/handler"
	"demo/app/cron/internal/svc"

	"github.com/robfig/cron/v3"
	"github.com/tal-tech/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/internal-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	schedule := cron.New(cron.WithSeconds())
	defer schedule.Stop()

	handler.RegisterTasks(schedule, ctx)

	fmt.Printf("Starting schedule \n")
	schedule.Start()

	exit := make(chan bool)
	<-exit
}

package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"os"
	"os/signal"
	"syscall"
	"time"

	"demo/app/mq/message/internal/config"
	"demo/app/mq/message/internal/handler"
	"demo/app/mq/message/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/message-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	//注册job
	group := service.NewServiceGroup()
	handler.RegisterDq(ctx, group)

	//捕捉信号
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-ch
		logx.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Printf("stop group")
			group.Stop()
			logx.Info("job exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

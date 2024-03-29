package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"demo/app/gateway/api/internal/config"
	"demo/app/gateway/api/internal/handler"
	"demo/app/gateway/api/internal/nsq"
	"demo/app/gateway/api/internal/svc"
	"demo/lib/errorx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 全局中间件
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Info("global middleware")
			next(w, r)
		}
	})

	handler.RegisterHandlers(server, ctx)

	var err error
	if err = nsq.InitConsumer(ctx); err != nil {
		panic(err)
	}

	// 自定义错误处理
	httpx.SetErrorHandler(errorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func errorHandler(err error) (int, interface{}) {
	switch e := err.(type) {
	case *errorx.RespError:
		fmt.Println(e.Data())
		return e.Status, e.Data()
	default:
		return http.StatusInternalServerError, err
	}
}

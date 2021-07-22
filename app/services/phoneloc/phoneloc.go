package main

import (
	"flag"
	"fmt"

	"demo/app/services/phoneloc/internal/config"
	"demo/app/services/phoneloc/internal/server"
	"demo/app/services/phoneloc/internal/svc"
	"demo/app/services/phoneloc/phoneloc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/phoneloc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewPhoneLocServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		phoneloc.RegisterPhoneLocServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

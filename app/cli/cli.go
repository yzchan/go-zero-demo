package main

import (
	"log"
	"os"

	"demo/app/cli/internal/config"
	"demo/app/cli/internal/handler"
	"demo/app/cli/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "configFile",
				Aliases: []string{"f"}, // 别名，简写参数
				Value:   "config.yaml", // 默认加载当前目录下的config.yaml
				Usage:   "配置文件路径",
			},
		},
		Before: func(context *cli.Context) error {
			// 拿到配置文件 注册svc
			var c config.Config
			conf.MustLoad(context.String("f"), &c)
			ctx := svc.NewServiceContext(c)
			handler.RegisterHandlers(context.App, ctx)
			return nil
		},
		Commands: []*cli.Command{},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

package handler

import (
	"demo/app/cli/internal/logic"
	"demo/app/cli/internal/svc"
	"github.com/urfave/cli/v2"
)

func CreateHandler(ctx *svc.ServiceContext) cli.ActionFunc {
	return func(context *cli.Context) error {
		l := logic.NewCreateLogic(context, ctx)
		return l.Create()
	}
}

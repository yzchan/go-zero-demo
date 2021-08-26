package handler

import (
	"demo/app/cli/internal/logic"
	"demo/app/cli/internal/svc"
	"github.com/urfave/cli/v2"
)

func ListHandler(ctx *svc.ServiceContext) cli.ActionFunc {
	return func(context *cli.Context) error {
		l := logic.NewListLogic(context, ctx)
		return l.List()
	}
}

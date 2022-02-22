package logic

import (
	"demo/app/cli/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/urfave/cli/v2"
	"time"
)

type ListLogic struct {
	logx.Logger
	ctx    *cli.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx *cli.Context, svcCtx *svc.ServiceContext) ListLogic {
	return ListLogic{
		Logger: logx.WithContext(ctx.Context),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List() error {
	logx.Info(l.ctx.String("o"))

	bT := time.Now()
	// do something
	eT := time.Since(bT)
	logx.Info("Run time: ", eT)
	return nil
}

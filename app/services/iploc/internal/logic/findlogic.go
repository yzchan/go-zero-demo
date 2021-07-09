package logic

import (
	"context"

	"demo/app/services/iploc/internal/svc"
	"demo/app/services/iploc/iploc"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLogic {
	return &FindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindLogic) Find(in *iploc.IpAddr) (*iploc.IpResult, error) {
	textA, textB := l.svcCtx.IpLoc.Find(in.Ip)
	return &iploc.IpResult{
		A: textA,
		B: textB,
	}, nil
}

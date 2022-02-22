package logic

import (
	"context"

	"demo/app/services/phoneloc/internal/svc"
	"demo/app/services/phoneloc/phoneloc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *FindLogic) Find(in *phoneloc.PhoneLocReq) (*phoneloc.PhoneLocResp, error) {
	// todo: add your logic here and delete this line

	return &phoneloc.PhoneLocResp{}, nil
}

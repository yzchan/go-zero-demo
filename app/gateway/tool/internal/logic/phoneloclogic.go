package logic

import (
	"context"

	"demo/app/gateway/tool/internal/svc"
	"demo/app/gateway/tool/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneLocLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneLocLogic(ctx context.Context, svcCtx *svc.ServiceContext) PhoneLocLogic {
	return PhoneLocLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PhoneLocLogic) PhoneLoc(req types.PhoneReq) (*types.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &types.CommonResp{}, nil
}

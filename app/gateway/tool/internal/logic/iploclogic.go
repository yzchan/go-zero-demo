package logic

import (
	"context"
	"demo/app/services/iploc/iploc"

	"demo/app/gateway/tool/internal/svc"
	"demo/app/gateway/tool/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IpLocLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIpLocLogic(ctx context.Context, svcCtx *svc.ServiceContext) IpLocLogic {
	return IpLocLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IpLocLogic) IpLoc(req types.IpReq) (ret *types.CommonResp, err error) {
	var resp *iploc.IpResult
	if resp, err = l.svcCtx.IpLoc.Find(l.ctx, &iploc.IpAddr{
		Ip: req.Ip,
	}); err != nil {
		return nil, err
	}

	return &types.CommonResp{
		Data: resp,
	}, nil
}

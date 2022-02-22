package handler

import (
	"context"
	"demo/app/queue/message/internal/logic"
	"demo/app/queue/message/internal/svc"
	"github.com/zeromicro/go-zero/core/service"
)

func RegisterDq(serverCtx *svc.ServiceContext, group *service.ServiceGroup) {
	//group.Add(logic.NewProducerLogic(context.Background(),serverCtx))
	group.Add(logic.NewConsumerLogic(context.Background(), serverCtx))
	group.Start()
}

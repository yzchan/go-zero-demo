package handler

import (
	"demo/app/cron/internal/logic"
	"demo/app/cron/internal/svc"
	"github.com/robfig/cron/v3"
)

func RegisterTasks(c *cron.Cron, ctx *svc.ServiceContext) {
	pc := logic.DemoTask{
		Ctx: ctx,
	}

	c.AddJob("* * * * * *", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(pc))
}

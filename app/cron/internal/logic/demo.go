package logic

import (
	"demo/app/cron/internal/svc"
	"fmt"
)

type DemoTask struct {
	Ctx *svc.ServiceContext
}

func (p DemoTask) Run() {
	fmt.Println("cron:xxxxxxxxxxxxxx")
}

package handlers

import (
	"github.com/nsqio/go-nsq"
	"github.com/zeromicro/go-zero/core/logx"
)

type Worker struct {
}

func NewWorkerHandler() *Worker {
	return &Worker{}
}

func (f *Worker) HandleMessage(msg *nsq.Message) (err error) {
	logx.Infof("msg body = %s\n", string(msg.Body))
	return
}

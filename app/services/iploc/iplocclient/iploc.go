// Code generated by goctl. DO NOT EDIT!
// Source: iploc.proto

//go:generate mockgen -destination ./iploc_mock.go -package iplocclient -source $GOFILE

package iplocclient

import (
	"context"

	"demo/app/services/iploc/iploc"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	IpAddr   = iploc.IpAddr
	IpResult = iploc.IpResult

	IpLoc interface {
		Find(ctx context.Context, in *IpAddr) (*IpResult, error)
	}

	defaultIpLoc struct {
		cli zrpc.Client
	}
)

func NewIpLoc(cli zrpc.Client) IpLoc {
	return &defaultIpLoc{
		cli: cli,
	}
}

func (m *defaultIpLoc) Find(ctx context.Context, in *IpAddr) (*IpResult, error) {
	client := iploc.NewIpLocClient(m.cli.Conn())
	return client.Find(ctx, in)
}

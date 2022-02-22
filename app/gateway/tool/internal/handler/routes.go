// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"demo/app/gateway/tool/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Example},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/iploc/:ipaddr",
					Handler: IpLocHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/phoneloc/:phone",
					Handler: PhoneLocHandler(serverCtx),
				},
			}...,
		),
	)
}

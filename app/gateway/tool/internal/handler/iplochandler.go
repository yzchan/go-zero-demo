package handler

import (
	"net/http"

	"demo/app/gateway/tool/internal/logic"
	"demo/app/gateway/tool/internal/svc"
	"demo/app/gateway/tool/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func IpLocHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IpReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewIpLocLogic(r.Context(), ctx)
		resp, err := l.IpLoc(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

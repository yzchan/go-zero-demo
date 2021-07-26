package handler

import (
	"net/http"

	"demo/app/gateway/tool/internal/logic"
	"demo/app/gateway/tool/internal/svc"
	"demo/app/gateway/tool/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func PhoneLocHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PhoneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPhoneLocLogic(r.Context(), ctx)
		resp, err := l.PhoneLoc(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

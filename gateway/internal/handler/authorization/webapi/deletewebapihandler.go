package webapi

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infra/gateway/internal/logic/authorization/webapi"
	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"
)

func DeleteWebApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteWebApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := webapi.NewDeleteWebApiLogic(r.Context(), svcCtx)
		resp, err := l.DeleteWebApi(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
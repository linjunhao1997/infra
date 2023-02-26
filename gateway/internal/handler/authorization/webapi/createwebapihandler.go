package webapi

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infra/gateway/internal/logic/authorization/webapi"
	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"
)

func CreateWebApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateWebApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := webapi.NewCreateWebApiLogic(r.Context(), svcCtx)
		resp, err := l.CreateWebApi(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

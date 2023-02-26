package webapi

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infra/gateway/internal/logic/authorization/webapi"
	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"
)

func QueryWebApiListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryWebApiListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := webapi.NewQueryWebApiListLogic(r.Context(), svcCtx)
		resp, err := l.QueryWebApiList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

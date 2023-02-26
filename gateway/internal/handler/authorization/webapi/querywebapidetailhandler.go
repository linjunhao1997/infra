package webapi

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infra/gateway/internal/logic/authorization/webapi"
	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"
)

func QueryWebApiDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryWebApiDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := webapi.NewQueryWebApiDetailLogic(r.Context(), svcCtx)
		resp, err := l.QueryWebApiDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

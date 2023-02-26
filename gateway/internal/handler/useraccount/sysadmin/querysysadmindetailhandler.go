package sysadmin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infra/gateway/internal/logic/useraccount/sysadmin"
	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"
)

func QuerySysadminDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QuerySysadminDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := sysadmin.NewQuerySysadminDetailLogic(r.Context(), svcCtx)
		resp, err := l.QuerySysadminDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

package sysadmin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infra/gateway/internal/logic/useraccount/sysadmin"
	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"
)

func QuerySysadminPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QuerySysadminPageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := sysadmin.NewQuerySysadminPageLogic(r.Context(), svcCtx)
		resp, err := l.QuerySysadminPage(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

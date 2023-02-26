package sysadmin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infra/gateway/internal/logic/useraccount/sysadmin"
	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"
)

func AuthenticateSysadminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthenticateSysadmin
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := sysadmin.NewAuthenticateSysadminLogic(r.Context(), svcCtx)
		resp, err := l.AuthenticateSysadmin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

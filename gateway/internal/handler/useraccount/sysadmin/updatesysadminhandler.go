package sysadmin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infra/gateway/internal/logic/useraccount/sysadmin"
	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"
)

func UpdateSysadminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSysadminReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := sysadmin.NewUpdateSysadminLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSysadmin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

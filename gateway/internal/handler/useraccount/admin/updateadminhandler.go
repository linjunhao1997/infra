package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infra/gateway/internal/logic/useraccount/admin"
	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"
)

func UpdateAdminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateAdminReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := admin.NewUpdateAdminLogic(r.Context(), svcCtx)
		resp, err := l.UpdateAdmin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infra/gateway/internal/logic/useraccount/admin"
	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"
)

func QueryAdminDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryAdminDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := admin.NewQueryAdminDetailLogic(r.Context(), svcCtx)
		resp, err := l.QueryAdminDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

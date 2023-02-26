package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infra/gateway/internal/logic/useraccount/admin"
	"infra/gateway/internal/svc"
	"infra/gateway/internal/types"
)

func QueryAdminListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryAdminListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := admin.NewQueryAdminListLogic(r.Context(), svcCtx)
		resp, err := l.QueryAdminList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

package authentication

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"infra/apiserver/internal/logic/authentication"
	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"
)

func GetAuthTokenByInternalHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InternalAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := authentication.NewGetAuthTokenByInternalLogic(r.Context(), svcCtx)
		resp, err := l.GetAuthTokenByInternal(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

package authentication

import (
	"net/http"
	"time"

	"infra/apiserver/internal/logic/authentication"
	"infra/apiserver/internal/svc"
	"infra/apiserver/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAuthTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAuthTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := authentication.NewGetAuthTokenLogic(r.Context(), svcCtx)
		resp, err := l.GetAuthToken(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			c := http.Cookie{
				Name:     "authToken",
				Path:     "/",
				Value:    resp.Data.AuthToken,
				HttpOnly: true,
				Expires:  time.Now().Add(30 * time.Minute),
			}
			w.Header().Add("Set-Cookie", c.String())
			httpx.OkJson(w, resp)
		}
	}
}

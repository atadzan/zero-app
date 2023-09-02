package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-app/internal/logic/auth"
	"zero-app/internal/svc"
	"zero-app/internal/types"
)

func SignInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignInReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewSignInLogic(r.Context(), svcCtx)
		resp, err := l.SignIn(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

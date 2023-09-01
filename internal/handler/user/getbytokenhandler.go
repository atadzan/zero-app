package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-app/internal/logic/user"
	"zero-app/internal/svc"
)

func GetByTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetByTokenLogic(r.Context(), svcCtx)
		resp, err := l.GetByToken()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

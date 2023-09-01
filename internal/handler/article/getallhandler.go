package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-app/internal/logic/article"
	"zero-app/internal/svc"
)

func GetAllHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := article.NewGetAllLogic(r.Context(), svcCtx)
		resp, err := l.GetAll()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

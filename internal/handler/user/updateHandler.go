package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-app/internal/svc"
	"zero-app/internal/types"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProfileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		//l := user.NewUpdateLogic(r.Context(), svcCtx)
		//resp, err := l.Update(&req)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}

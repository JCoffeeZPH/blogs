package handler

import (
	"blogs/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"blogs/app/admin/api/internal/logic"
	"blogs/app/admin/api/internal/svc"
	"blogs/app/admin/api/internal/types"
)

func AreaDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAreaDataRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAreaDataLogic(r.Context(), svcCtx)
		resp, err := l.AreaData(&req)
		response.Response(w, resp, err)
	}
}

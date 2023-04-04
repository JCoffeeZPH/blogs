package core

import (
	"blogs/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"blogs/app/core/api/internal/logic/core"
	"blogs/app/core/api/internal/svc"
	"blogs/app/core/api/internal/types"
)

func GetAllArticlesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllArticlesRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := core.NewGetAllArticlesLogic(r.Context(), svcCtx)
		resp, err := l.GetAllArticles(&req)
		response.Response(w, resp, err)
	}
}

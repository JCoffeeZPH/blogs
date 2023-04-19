package handler

import (
	"blogs/common/response"
	"net/http"

	"blogs/app/admin/api/internal/logic"
	"blogs/app/admin/api/internal/svc"
)

func SearchTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSearchTagLogic(r.Context(), svcCtx)
		resp, err := l.SearchTag()
		response.Response(w, resp, err)
	}
}

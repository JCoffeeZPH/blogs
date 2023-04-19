package handler

import (
	"blogs/common/response"
	"net/http"

	"blogs/app/admin/api/internal/logic"
	"blogs/app/admin/api/internal/svc"
)

func SearchCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSearchCategoryLogic(r.Context(), svcCtx)
		resp, err := l.SearchCategory()
		response.Response(w, resp, err)
	}
}

package core

import (
	"blogs/common/response"
	"net/http"

	"blogs/app/core/api/internal/logic/core"
	"blogs/app/core/api/internal/svc"
)

func ApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := core.NewApiLogic(r.Context(), svcCtx)
		resp, err := l.Api()
		response.Response(w, resp, err)
	}
}

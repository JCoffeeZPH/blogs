package core

import (
	"blogs/common/response"
	"net/http"

	"blogs/app/core/api/internal/logic/core"
	"blogs/app/core/api/internal/svc"
)

func ReportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := core.NewReportLogic(r.Context(), svcCtx)

		resp, err := l.Report(r)
		response.Response(w, resp, err)
	}
}

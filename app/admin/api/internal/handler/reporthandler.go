package handler

import (
	"blogs/common/response"
	"net/http"

	"blogs/app/admin/api/internal/logic"
	"blogs/app/admin/api/internal/svc"
)

func ReportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewReportLogic(r.Context(), svcCtx)
		err := l.Report(r)
		response.Response(w, nil, err)
	}
}

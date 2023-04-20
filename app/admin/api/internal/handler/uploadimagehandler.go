package handler

import (
	"blogs/app/admin/api/internal/logic"
	"blogs/app/admin/api/internal/svc"
	"blogs/common/response"
	"net/http"
)

func UploadImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadImageLogic(r.Context(), svcCtx)
		fileUrl, err := l.UploadImage(r)
		response.Response(w, fileUrl, err)
	}
}

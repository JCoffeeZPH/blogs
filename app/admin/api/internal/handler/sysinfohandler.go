package handler

import (
	"blogs/common/constants"
	"blogs/common/response"
	"net/http"
	"strconv"

	"blogs/app/admin/api/internal/logic"
	"blogs/app/admin/api/internal/svc"
)

func SysInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSysInfoLogic(r.Context(), svcCtx)

		userId, _ := strconv.ParseInt(r.Header.Get(constants.UserId), 10, 64)
		resp, err := l.SysInfo(userId)
		response.Response(w, resp, err)
	}
}

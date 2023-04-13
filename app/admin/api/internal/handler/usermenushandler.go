package handler

import (
	"blogs/common/constants"
	"blogs/common/response"
	sess "blogs/models/session"
	"net/http"

	"blogs/app/admin/api/internal/logic"
	"blogs/app/admin/api/internal/svc"
)

func UserMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserMenusLogic(r.Context(), svcCtx)

		userInfoId := sess.GetInt64(r, constants.UserId)
		resp, err := l.UserMenus(userInfoId)

		response.Response(w, resp, err)
	}
}

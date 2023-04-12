package handler

import (
	"blogs/common/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"

	"blogs/app/admin/api/internal/logic"
	"blogs/app/admin/api/internal/svc"
	"blogs/app/admin/api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		jwtToken := getJwtToken(r)
		resp, err := l.Login(&req, jwtToken)
		response.Response(w, resp, err)
	}
}

func getJwtToken(r *http.Request) string {
	header := r.Header.Get("Authorization")
	splits := strings.Split(header, " ")
	if len(splits) == 2 && strings.ToLower(splits[0]) == "bearer" {
		return splits[1]
	} else {
		return ""
	}
}

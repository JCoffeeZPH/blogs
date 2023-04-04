package core

import (
	"blogs/app/core/api/internal/types"
	"blogs/common/response"
	"net/http"

	"blogs/app/core/api/internal/logic/core"
	"blogs/app/core/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPhotoAlbumsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPhotoAlbumsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := core.NewGetPhotoAlbumsLogic(r.Context(), svcCtx)
		resp, err := l.GetPhotoAlbums(&req)
		response.Response(w, resp, err)
	}
}

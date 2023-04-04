package core

import (
	"blogs/common/response"
	"net/http"

	"blogs/app/core/api/internal/logic/core"
	"blogs/app/core/api/internal/svc"
)

func CategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := core.NewCategoryLogic(r.Context(), svcCtx)
		resp, err := l.Category()
		response.Response(w, resp, err)
	}
}

type T struct {
	Flag    bool   `json:"flag"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Records []struct {
			Id             int    `json:"id"`
			ArticleCover   string `json:"articleCover"`
			ArticleTitle   string `json:"articleTitle"`
			ArticleContent string `json:"articleContent"`
			IsTop          int    `json:"isTop"`
			IsFeatured     int    `json:"isFeatured"`
			CategoryName   string `json:"categoryName"`
			Status         int    `json:"status"`
			CreateTime     string `json:"createTime"`
			UpdateTime     string `json:"updateTime"`
		} `json:"records"`
		Count int `json:"count"`
	} `json:"data"`
}

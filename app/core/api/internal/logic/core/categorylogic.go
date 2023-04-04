package core

import (
	"context"

	"blogs/app/core/api/internal/svc"
	"blogs/app/core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryLogic {
	return &CategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryLogic) Category() (*types.GetAllCategoriesResponse, error) {
	categories := l.svcCtx.CategoryDao.GetAllCategories()
	resp := &types.GetAllCategoriesResponse{}
	for _, category := range categories {
		resp.CategoryDetails = append(resp.CategoryDetails, types.CategoryDetail{
			Id:           category.CategoryId,
			CategoryName: category.CategoryName,
			ArticleCount: category.ArticleCount,
		})
	}
	return resp, nil
}

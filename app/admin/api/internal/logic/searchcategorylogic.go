package logic

import (
	"context"

	"blogs/app/admin/api/internal/svc"
	"blogs/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCategoryLogic {
	return &SearchCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchCategoryLogic) SearchCategory() ([]*types.SearchCategoryResponse, error) {
	var resp []*types.SearchCategoryResponse
	categories := l.svcCtx.CategoryDao.GetAllCategories()
	for _, category := range categories {
		resp = append(resp, &types.SearchCategoryResponse{
			Id:           category.CategoryId,
			CategoryName: category.CategoryName,
		})
	}
	return resp, nil
}

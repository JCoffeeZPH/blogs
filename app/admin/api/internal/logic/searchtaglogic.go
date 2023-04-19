package logic

import (
	"context"

	"blogs/app/admin/api/internal/svc"
	"blogs/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTagLogic {
	return &SearchTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchTagLogic) SearchTag() ([]*types.SearchTagResponse, error) {
	var resp []*types.SearchTagResponse
	tags := l.svcCtx.TagDao.GetAllTags()

	for _, tag := range tags {
		resp = append(resp, &types.SearchTagResponse{
			Id:      tag.Id,
			TagName: tag.TagName,
		})
	}
	return resp, nil
}

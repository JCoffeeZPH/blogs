package logic

import (
	"blogs/common/constants"
	"blogs/lib/cache"
	"context"

	"blogs/app/admin/api/internal/svc"
	"blogs/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysInfoLogic {
	return &SysInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysInfoLogic) SysInfo(userId int64) (*types.GetSysInfoResponse, error) {
	resp := &types.GetSysInfoResponse{}

	viewCount, err := cache.GetCount(constants.BlogsViewCountKey)
	if err != nil {
		logx.Errorf("get view count from redis failed, err: %+v", err)
		return nil, err
	}

	commentsCount := l.svcCtx.CommentDao.CountComments(map[string]interface{}{"type": 2})
	loginUsers := l.svcCtx.UserInfoDao.CountUser(nil)
	articleCount := l.svcCtx.ArticleDao.Count(map[string]interface{}{"is_delete": constants.ArticleStatusNotDelete})

	resp.ViewCount = viewCount

	return nil, nil
}

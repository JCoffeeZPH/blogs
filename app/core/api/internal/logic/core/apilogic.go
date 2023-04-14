package core

import (
	"blogs/common/constants"
	"blogs/lib/cache"
	"context"
	"encoding/json"

	"blogs/app/core/api/internal/svc"
	"blogs/app/core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiLogic {
	return &ApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiLogic) Api() (*types.ApiResponse, error) {
	resp := &types.ApiResponse{}
	articleCount := l.svcCtx.ArticleDao.Count(map[string]interface{}{"is_delete": constants.ArticleStatusNotDelete})
	tagCount := l.svcCtx.TagDao.Count()
	categoryCount := l.svcCtx.CategoryDao.Count()
	talkCount := l.svcCtx.TalkDao.Count()

	viewCount, err := cache.GetCount(constants.BlogsViewCountKey)
	if err != nil {
		logx.Errorf("get view count from redis failed, err: %+v", err)
	}

	siteConf, err := cache.GetWebsiteConfig(constants.DefaultWebsiteConfigKey)
	if err != nil || len(siteConf) == 0 {
		logx.Errorf("get website config from redis failed, err: %+v", err)
		siteConf = l.svcCtx.WebsiteConfigDao.GetWebsiteConfig().Config
		go cache.SetByKey(constants.DefaultWebsiteConfigKey, siteConf)
	}

	resp.ArticleCount = articleCount
	resp.TagCount = tagCount
	resp.CategoryCount = categoryCount
	resp.TalkCount = talkCount
	resp.ViewCount = viewCount

	websiteConfig := types.WebSiteConfig{}

	json.Unmarshal([]byte(siteConf), &websiteConfig)
	resp.WebSiteConfig = websiteConfig

	return resp, nil
}

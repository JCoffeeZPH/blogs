package logic

import (
	"blogs/app/admin/api/internal/svc"
	"blogs/app/admin/api/internal/types"
	"blogs/common/constants"
	"blogs/common/utils"
	"blogs/lib/cache"
	"context"

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

	messageCount := l.svcCtx.CommentDao.CountComments(map[string]interface{}{"type": 2})
	loginUsers := l.svcCtx.UserInfoDao.CountUser(nil)
	articleCount := l.svcCtx.ArticleDao.Count(map[string]interface{}{"is_delete": constants.ArticleStatusNotDelete})
	uniqueViews := l.svcCtx.UniqueDao.GetUniqueByTime(utils.GetLast7DayTime())
	tags := l.svcCtx.TagDao.GetAllTags()
	categories := l.svcCtx.CategoryDao.GetAllCategories()
	articleDateStat := l.svcCtx.ArticleDao.GetArticleStatByDate()

	resp.ViewsCount = viewCount
	resp.ArticleCount = articleCount
	resp.MessageCount = messageCount
	resp.UserCount = loginUsers
	resp.ArticleRanks = l.getTopArticlesCount()
	for _, uniqueView := range uniqueViews {
		resp.UniqueViews = append(resp.UniqueViews, types.UniqueView{
			Day:        utils.TimeFormatDay(uniqueView.Ctime),
			ViewsCount: uniqueView.ViewsCount,
		})
	}

	for _, tag := range tags {
		resp.Tags = append(resp.Tags, types.Tag{
			Id:      tag.Id,
			TagName: tag.TagName,
		})
	}

	for _, category := range categories {
		resp.Categories = append(resp.Categories, types.Category{
			Id:           category.CategoryId,
			CategoryName: category.CategoryName,
			ArticleCount: category.ArticleCount,
		})
	}

	for _, stat := range articleDateStat {
		resp.ArticleStats = append(resp.ArticleStats, types.ArticleStat{
			Date:  stat.Date,
			Count: stat.Count,
		})
	}

	return resp, nil
}

func (l *SysInfoLogic) getTopArticlesCount() []types.ArticleRank {
	cacheArticles := cache.ZRevRangeWithScores(constants.ArticleRankKey, 0, 4)
	articleMap := make(map[int64]int64)
	articleIds := make([]int64, 0)
	for _, article := range cacheArticles {
		articleId, _ := article.Member.(int)
		articleMap[int64(articleId)] = int64(article.Score)
		articleIds = append(articleIds, int64(articleId))
	}

	topArticles := l.svcCtx.ArticleDao.GetArticlesById(articleIds)

	articleRanks := make([]types.ArticleRank, 0)
	for _, article := range topArticles {
		articleRanks = append(articleRanks, types.ArticleRank{
			ArticleTitle: article.ArticleTitle,
			ViewsCount:   articleMap[article.Id],
		})
	}
	return articleRanks
}

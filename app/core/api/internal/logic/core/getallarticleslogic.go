package core

import (
	"blogs/common/utils"
	"blogs/models"
	"context"

	"blogs/app/core/api/internal/svc"
	"blogs/app/core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllArticlesLogic {
	return &GetAllArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetAllArticles logic
// 1、get articles from article_tab and take out articleIds
// 2、get tagIds from article_tag_tab according article_ids
// 3、get tags from tag_tab by tagIds
func (l *GetAllArticlesLogic) GetAllArticles(req *types.GetAllArticlesRequest) (*types.GetAllArticlesResponse, error) {
	articles := l.svcCtx.ArticleDao.GetAllArticles(int(req.Current), int(req.Size))
	articleIds := make([]int64, 0)
	for _, article := range articles {
		articleIds = append(articleIds, article.Id)
	}

	count := len(articles)

	articlesTags := l.svcCtx.ArticleTagDao.GetTagIdsByArticleIds(articleIds)
	articleTagIdMap := make(map[int64][]int64)
	tagsIds := make([]int64, 0)
	for _, tag := range articlesTags {
		if _, exist := articleTagIdMap[tag.ArticleId]; !exist {
			articleTagIdMap[tag.ArticleId] = make([]int64, 0)
		}
		articleTagIdMap[tag.ArticleId] = append(articleTagIdMap[tag.ArticleId], tag.TagId)
		tagsIds = append(tagsIds, tag.TagId)
	}

	tags := l.svcCtx.TagDao.GetTagsByIds(tagsIds)
	tagMap := make(map[int64]models.Tag)
	for _, tag := range tags {
		tagMap[tag.Id] = tag
	}

	userInfo := l.svcCtx.UserInfoDao.GetUserInfoById(1)

	resp := &types.GetAllArticlesResponse{Count: int32(count)}
	for _, article := range articles {
		articleTags := make([]types.Tag, 0)
		tagIds := articleTagIdMap[article.Id]

		for _, tagId := range tagIds {
			tag := tagMap[tagId]
			articleTags = append(articleTags, types.Tag{
				Id:         tag.Id,
				TagName:    tag.TagName,
				CreateTime: utils.TimeFormat(tag.CreateTime),
				UpdateTime: utils.TimeFormat(tag.UpdateTime),
			})
		}

		resp.Records = append(resp.Records, types.Record{
			Id:           article.Id,
			ArticleCover: article.ArticleCover,
			ArticleTitle: article.ArticleTitle,
			IsTop:        article.IsTop,
			IsFeatured:   article.IsFeatured,
			CategoryName: article.CategoryName,
			Status:       article.Status,
			CreateTime:   utils.TimeFormat(article.CreateTime),
			UpdateTime:   utils.TimeFormat(article.UpdateTime),
			Tags:         articleTags,
			Author: types.Author{
				Id:          userInfo.Id,
				Email:       userInfo.Email,
				Nickname:    userInfo.Nickname,
				Avatar:      userInfo.Avatar,
				Intro:       userInfo.Intro,
				Website:     userInfo.Website,
				IsSubscribe: userInfo.IsSubscribe,
				IsDisable:   userInfo.IsDisable,
				CreateTime:  utils.TimeFormat(userInfo.CreateTime),
				UpdateTime:  utils.TimeFormat(userInfo.UpdateTime),
			},
		})
	}
	return resp, nil
}

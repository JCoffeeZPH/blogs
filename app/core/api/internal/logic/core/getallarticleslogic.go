package core

import (
	"blogs/common/utils"
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

func (l *GetAllArticlesLogic) GetAllArticles(req *types.GetAllArticlesRequest) (*types.GetAllArticlesResponse, error) {
	records := l.svcCtx.ArticleDao.GetAllArticles(int(req.Current), int(req.Size))
	resp := &types.GetAllArticlesResponse{}
	for _, record := range records {
		tags := make([]types.Tag, 0)
		for _, tag := range record.Tags {
			tags = append(tags, types.Tag{
				Id:         tag.Id,
				TagName:    tag.TagName,
				CreateTime: utils.TimeFormat(tag.CreateTime),
				UpdateTime: utils.TimeFormat(tag.UpdateTime),
			})
		}

		resp.Records = append(resp.Records, types.Record{
			Id:           record.Id,
			ArticleCover: record.ArticleCover,
			ArticleTitle: record.ArticleTitle,
			IsTop:        record.IsTop,
			IsFeatured:   record.IsFeatured,
			CategoryName: record.CategoryName,
			Status:       record.Status,
			CreateTime:   utils.TimeFormat(record.CreateTime),
			UpdateTime:   utils.TimeFormat(record.UpdateTime),
			Tags:         tags,
			Author: types.Author{
				Id:          record.Author.Id,
				Email:       record.Author.Email,
				Nickname:    record.Author.Nickname,
				Avatar:      record.Author.Avatar,
				Intro:       record.Author.Intro,
				Website:     record.Author.Website,
				IsSubscribe: record.Author.IsSubscribe,
				IsDisable:   record.Author.IsDisable,
				CreateTime:  utils.TimeFormat(record.Author.CreateTime),
				UpdateTime:  utils.TimeFormat(record.Author.UpdateTime),
			},
		})
	}
	return resp, nil
}

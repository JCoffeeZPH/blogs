package dao

import (
	"blogs/common/errorx"
	"blogs/models"
	"gorm.io/gorm"
)

type ArticleTagDao interface {
	Count() int64
	GetTagIdsByArticleIds(articleIds []int64) []models.ArticleTag
}

type ArticleTagDaoImpl struct {
	db *gorm.DB
}

func NewArticleTagDao() ArticleTagDao {
	return &ArticleTagDaoImpl{
		db: GetDBClient(),
	}
}

func (dao *ArticleTagDaoImpl) Count() int64 {
	var count int64
	err := dao.db.Model(&models.ArticleTag{}).Select("count(distinct(tag_id))").Count(&count).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return count
}

func (dao *ArticleTagDaoImpl) GetTagIdsByArticleIds(articleIds []int64) []models.ArticleTag {
	var articleTag []models.ArticleTag
	if err := dao.db.Where("article_id in (?)", articleIds).Find(&articleTag).Error; err != nil {
		panic(errorx.DBError{Err: err})
	}

	return articleTag
}

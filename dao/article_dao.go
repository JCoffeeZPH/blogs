package dao

import (
	"blogs/common/errorx"
	"blogs/models"
	"gorm.io/gorm"
)

type ArticleDao interface {
	Count() int64
	GetAllArticles(current, size int) []models.ArticleWithCategory
}

type ArticleDaoImpl struct {
	db *gorm.DB
}

func NewArticleDao() ArticleDao {
	return &ArticleDaoImpl{
		db: GetDBClient(),
	}
}

func (dao *ArticleDaoImpl) Count() int64 {
	var count int64
	err := dao.db.Model(&models.Article{}).Where("is_delete = ?", 0).Count(&count).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return count
}

func (dao *ArticleDaoImpl) GetAllArticles(current, size int) []models.ArticleWithCategory {
	var records []models.ArticleWithCategory

	err := dao.db.Table("article_tab article").Select("article.id, "+
		"article.user_id, "+
		"article.category_id,"+
		"article.article_cover, "+
		"article.article_title, "+
		"article.article_content, "+
		"article.is_top, "+
		"article.is_featured, "+
		"article.is_delete, "+
		"article.status, "+
		"article.create_time, "+
		"article.update_time,"+
		"category.category_name").
		Where("user_id = ? and is_delete = ? and status in (1, 2)", 1, 0).Order("id desc").Limit(size).Offset(current).
		Joins("left join category_tab category on article.category_id = category.id").Find(&records).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}

	return records
}

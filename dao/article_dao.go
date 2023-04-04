package dao

import (
	"blogs/common/errorx"
	"blogs/models"
	"gorm.io/gorm"
)

type ArticleDao interface {
	Count() int64
	GetAllArticles(current, size int) []models.ArticleCard
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

func (dao *ArticleDaoImpl) GetAllArticles(current, size int) []models.ArticleCard {
	var records []models.ArticleCard

	subQuery := dao.db.Table("article_tab").Select("id, "+
		"user_id, "+
		"category_id,"+
		" article_cover, "+
		"article_title, "+
		"article_content, "+
		"is_top, "+
		"is_featured, "+
		"is_delete, "+
		"status, "+
		"create_time, "+
		"update_time").
		Where("is_delete = ? and status in (?, ?)", 0, 1, 2).Order("id desc").Limit(size).Offset(current)

	err := dao.db.Table("(?) as article", subQuery).
		Select("article.id as id, " +
			"article_cover, " +
			"article_title, " +
			"SUBSTR(article_content, 1, 500) AS article_content," +
			"is_top," +
			"is_featured," +
			"status," +
			"article.create_time as create_time," +
			"article.update_time as update_time," +
			"user.nickname as author_nickname," +
			"user.website as author_website," +
			"user.avatar as author_avatar," +
			"category.category_name as category_name," +
			"tag_name").Joins("left join article_tag_tab article_tag on article.id = article_tag.article_id").
		Joins("left join tag_tab tag on tag.id = article_tag.tag_id").
		Joins("left join category_tab category on article.category_id = category.id").
		Joins("left join user_info_tab user on article.user_id = user.id").Scan(&records)

	if err != nil {
		panic(errorx.DBError{Err: err})
	}

	return records
}

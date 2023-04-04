package dao

import (
	"blogs/common/errorx"
	"blogs/models"
	"github.com/jinzhu/gorm"
)

type ArticleDao interface {
	Count() int64
}

type ArticleDaoImpl struct {
	db *gorm.DB
}

func NewArticleDao() ArticleDao {
	return &ArticleDaoImpl{
		db: GetDBClient(),
	}
}

func (dao ArticleDaoImpl) Count() int64 {
	var count int64
	err := dao.db.Model(&models.Article{}).Where("is_delete = ?", 0).Count(&count).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return count
}

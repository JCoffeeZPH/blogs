package dao

import (
	"blogs/common/errorx"
	"blogs/models"
	"github.com/jinzhu/gorm"
)

type ArticleTagDao interface {
	Count() int64
}

type ArticleTagDaoImpl struct {
	db *gorm.DB
}

func NewArticleTagDao() ArticleTagDao {
	return &ArticleTagDaoImpl{
		db: GetDBClient(),
	}
}

func (dao ArticleTagDaoImpl) Count() int64 {
	var count int64
	err := dao.db.Model(&models.ArticleTag{}).Select("count(distinct(tag_id))").Count(&count).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return count
}

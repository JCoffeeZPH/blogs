package dao

import (
	"blogs/common/errorx"
	"blogs/models/db"
	"gorm.io/gorm"
)

type CommentDao interface {
	CountComments(param map[string]interface{}) int64
}

type CommentDaoImpl struct {
	db *gorm.DB
}

func NewCommentDao() CommentDao {
	return &CommentDaoImpl{
		db: GetDBClient(),
	}
}

func (dao *CommentDaoImpl) CountComments(param map[string]interface{}) int64 {
	var count int64
	if err := dao.db.Model(&db.Comment{}).Where(param).Count(&count).Error; err != nil {
		panic(errorx.DBError{Err: err})
	}
	return count
}

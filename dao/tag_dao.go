package dao

import (
	"blogs/common/errorx"
	"blogs/models"
	"github.com/jinzhu/gorm"
)

type TagDao interface {
	Count() int64
}

type TagDaoImpl struct {
	db *gorm.DB
}

func NewTagDao() TagDao {
	return &TagDaoImpl{
		db: GetDBClient(),
	}
}

func (dao TagDaoImpl) Count() int64 {
	var count int64
	err := dao.db.Model(&models.Tag{}).Count(&count).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return count
}

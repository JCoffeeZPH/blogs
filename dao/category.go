package dao

import (
	"blogs/common/errorx"
	"blogs/models"
	"github.com/jinzhu/gorm"
)

type CategoryDao interface {
	Count() int64
}

type CategoryDaoImpl struct {
	db *gorm.DB
}

func NewCategoryDao() CategoryDao {
	return &CategoryDaoImpl{
		db: GetDBClient(),
	}
}

func (dao CategoryDaoImpl) Count() int64 {
	var count int64
	err := dao.db.Model(&models.Category{}).Count(&count).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return count
}

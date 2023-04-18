package dao

import (
	"blogs/common/errorx"
	"blogs/models/db"
	"gorm.io/gorm"
)

type UniqueViewDao interface {
	GetUniqueByTime(timestamps []uint32) []db.UniqueView
}

type UniqueViewDaoImpl struct {
	db *gorm.DB
}

func NewUniqueViewDao() UniqueViewDao {
	return &UniqueViewDaoImpl{
		db: GetDBClient(),
	}
}

func (dao UniqueViewDaoImpl) GetUniqueByTime(timestamps []uint32) []db.UniqueView {
	views := make([]db.UniqueView, 0)
	if err := dao.db.Model(&db.UniqueView{}).Where("ctime in (?)", timestamps).Order("id").Find(&views).Error; err != nil {
		panic(errorx.DBError{Err: err})
	}
	return views
}

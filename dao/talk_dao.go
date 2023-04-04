package dao

import (
	"blogs/common/errorx"
	"blogs/models"
	"gorm.io/gorm"
)

type TalkDao interface {
	Count() int64
}

type TalkDaoImpl struct {
	db *gorm.DB
}

func NewTalkDao() TalkDao {
	return &TalkDaoImpl{
		db: GetDBClient(),
	}
}

func (dao *TalkDaoImpl) Count() int64 {
	var count int64
	err := dao.db.Model(&models.Talk{}).Count(&count).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return count
}

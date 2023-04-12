package dao

import (
	"blogs/common/errorx"
	"blogs/models/db"
	"gorm.io/gorm"
)

type UserInfoDao interface {
	GetUserInfoById(userId int64) db.UserInfo
}

type UserInfoDaoImpl struct {
	db *gorm.DB
}

func NewUserInfoDaoDao() UserInfoDao {
	return &UserInfoDaoImpl{
		db: GetDBClient(),
	}
}

func (dao *UserInfoDaoImpl) GetUserInfoById(userId int64) db.UserInfo {
	userInfo := db.UserInfo{}
	if err := dao.db.Where("id = ?", userId).First(&userInfo).Error; err != nil {
		panic(errorx.DBError{Err: err})
	}

	return userInfo
}

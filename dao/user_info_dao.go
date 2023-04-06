package dao

import (
	"blogs/common/errorx"
	"blogs/models"
	"gorm.io/gorm"
)

type UserInfoDao interface {
	GetUserInfoById(userId int64) models.UserInfo
}

type UserInfoDaoImpl struct {
	db *gorm.DB
}

func NewUserInfoDaoDao() UserInfoDao {
	return &UserInfoDaoImpl{
		db: GetDBClient(),
	}
}

func (dao *UserInfoDaoImpl) GetUserInfoById(userId int64) models.UserInfo {
	userInfo := models.UserInfo{}
	if err := dao.db.Where("id = ?", userId).First(&userInfo).Error; err != nil {
		panic(errorx.DBError{Err: err})
	}

	return userInfo
}

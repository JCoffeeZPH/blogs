package dao

import (
	"blogs/common/errorx"
	"blogs/models/db"
	"gorm.io/gorm"
)

type UserInfoDao interface {
	GetUserInfoById(userId int64) db.UserInfo
	CountUser(param map[string]interface{}) int64
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

func (dao *UserInfoDaoImpl) CountUser(param map[string]interface{}) int64 {
	var count int64
	err := dao.db.Model(&db.UserInfo{}).Where(param).Count(&count).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return count
}

package dao

import (
	"blogs/common/errorx"
	"blogs/models/db"
	"gorm.io/gorm"
)

type UserAuthDao interface {
	GetUserInfo(username, password string) *db.UserDetails
	Update(id int64, params map[string]interface{})
	CountUser(param map[string]interface{}) int64
}

type UserAuthDaoImpl struct {
	db *gorm.DB
}

func NewUserAuthDao() UserAuthDao {
	return &UserAuthDaoImpl{
		db: GetDBClient(),
	}
}

func (dao *UserAuthDaoImpl) GetUserInfo(username, password string) *db.UserDetails {
	var userDetails db.UserDetails
	err := dao.db.Table("user_auth_tab auth").Select("auth.user_info_id,"+
		"auth.username,"+
		"auth.login_type,"+
		"auth.ip_address,"+
		"auth.ip_source,"+
		"auth.last_login_time,"+
		"info.email,"+
		"info.nickname,"+
		"info.avatar,"+
		"info.intro,"+
		"info.website,"+
		"info.is_subscribe,"+
		"info.is_disable,"+
		"role.role_id").
		Where("username = ? and password = ?", username, password).
		Joins("left join user_info_tab info on auth.user_info_id = info.id").
		Joins("left join user_role_tab role on info.id = role.user_id").First(&userDetails).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return &userDetails
}

func (dao *UserAuthDaoImpl) Update(id int64, params map[string]interface{}) {
	err := dao.db.Model(&db.UserAuth{}).Where("id = ?", id).Updates(params).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
}

func (dao *UserAuthDaoImpl) CountUser(param map[string]interface{}) int64 {
	var count int64
	err := dao.db.Model(&db.UserAuth{}).Where(param).Count(&count).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return count
}

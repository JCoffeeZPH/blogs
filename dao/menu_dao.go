package dao

import (
	"blogs/common/errorx"
	"blogs/models/db"
	"gorm.io/gorm"
)

type MenuDao interface {
	GetMenusByUserId(userId int64) []db.Menu
}

type MenuDaoImpl struct {
	db *gorm.DB
}

func NewMenuDao() MenuDao {
	return &MenuDaoImpl{
		db: GetDBClient(),
	}
}

func (dao *MenuDaoImpl) GetMenusByUserId(userId int64) []db.Menu {
	var menus []db.Menu

	if err := dao.db.Table("user_role_tab ur").Where("user_id = ?", userId).Select("m.id, m.name, m.path, m.component, m.icon, m.is_hidden, m.parent_id, m.order_num").
		Joins("join role_menu_tab rm ON ur.role_id = rm.role_id").
		Joins("join menu_tab m on rm.menu_id = m.id").Find(&menus).Error; err != nil {
		panic(errorx.DBError{Err: err})
	}
	return menus
}

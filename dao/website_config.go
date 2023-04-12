package dao

import (
	"blogs/common/errorx"
	"blogs/models/db"
	"gorm.io/gorm"
)

type WebsiteConfigDao interface {
	GetWebsiteConfig() db.WebsiteConfig
}

type WebsiteConfigDaoImpl struct {
	db *gorm.DB
}

func NewWebsiteConfigDao() WebsiteConfigDao {
	return &WebsiteConfigDaoImpl{
		db: GetDBClient(),
	}
}

func (dao *WebsiteConfigDaoImpl) GetWebsiteConfig() db.WebsiteConfig {
	var conf db.WebsiteConfig
	err := dao.db.Where("id = ?", 2).First(&conf).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return conf
}

package dao

import (
	"blogs/common/errorx"
	"blogs/models"
	"github.com/jinzhu/gorm"
)

type WebsiteConfigDao interface {
	GetWebsiteConfig() models.WebsiteConfig
}

type WebsiteConfigDaoImpl struct {
	db *gorm.DB
}

func NewWebsiteConfigDao() WebsiteConfigDao {
	return &WebsiteConfigDaoImpl{
		db: GetDBClient(),
	}
}

func (dao WebsiteConfigDaoImpl) GetWebsiteConfig() models.WebsiteConfig {
	var conf models.WebsiteConfig
	err := dao.db.Where("id = ?", 2).First(&conf).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return conf
}

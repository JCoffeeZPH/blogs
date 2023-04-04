package dao

import (
	"blogs/common/errorx"
	"blogs/models"
	"github.com/jinzhu/gorm"
)

type PhotoAlbumDao interface {
	GetPhotoAlbums(offset, limit int64) []models.PhotoAlbum
}

type PhotoAlbumDaoImpl struct {
	db *gorm.DB
}

func NewPhotoAlbumDao() PhotoAlbumDao {
	return &PhotoAlbumDaoImpl{
		db: GetDBClient(),
	}
}

func (dao PhotoAlbumDaoImpl) GetPhotoAlbums(offset, limit int64) []models.PhotoAlbum {
	var albums []models.PhotoAlbum
	err := dao.db.Where("is_delete = ? and status = ?", 0, 1).Offset(offset).Limit(limit).Find(&albums).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return albums
}

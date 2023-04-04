package dao

import (
	"blogs/common/errorx"
	"blogs/models"
	"gorm.io/gorm"
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

func (dao *PhotoAlbumDaoImpl) GetPhotoAlbums(offset, limit int64) []models.PhotoAlbum {
	var albums []models.PhotoAlbum
	err := dao.db.Where("is_delete = ? and status = ?", 0, 1).Offset(int(offset)).Limit(int(limit)).Find(&albums).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return albums
}

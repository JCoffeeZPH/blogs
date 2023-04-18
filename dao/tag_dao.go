package dao

import (
	"blogs/common/errorx"
	"blogs/models/db"
	"gorm.io/gorm"
)

type TagDao interface {
	Count() int64
	GetTagsByIds(tagIds []int64) []db.Tag
	GetAllTags() []db.Tag
}

type TagDaoImpl struct {
	db *gorm.DB
}

func NewTagDao() TagDao {
	return &TagDaoImpl{
		db: GetDBClient(),
	}
}

func (dao *TagDaoImpl) Count() int64 {
	var count int64
	err := dao.db.Model(&db.Tag{}).Count(&count).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return count
}

func (dao *TagDaoImpl) GetTagsByIds(tagIds []int64) []db.Tag {
	var tags []db.Tag
	if err := dao.db.Distinct().Where("id in (?)", tagIds).Find(&tags).Error; err != nil {
		panic(errorx.DBError{Err: err})
	}

	return tags
}

func (dao *TagDaoImpl) GetAllTags() []db.Tag {
	var tags []db.Tag
	if err := dao.db.Model(&db.Tag{}).Select("id, tag_name").Order("id").Find(&tags).Error; err != nil {
		panic(errorx.DBError{Err: err})
	}
	return tags
}

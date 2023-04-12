package dao

import (
	"blogs/common/errorx"
	"blogs/models/db"
	"gorm.io/gorm"
)

type CategoryDao interface {
	Count() int64
	GetAllCategories() []db.CategoryDetail
}

type CategoryDaoImpl struct {
	db *gorm.DB
}

func NewCategoryDao() CategoryDao {
	return &CategoryDaoImpl{
		db: GetDBClient(),
	}
}

func (dao *CategoryDaoImpl) Count() int64 {
	var count int64
	err := dao.db.Model(&db.Category{}).Count(&count).Error
	if err != nil {
		panic(errorx.DBError{Err: err})
	}
	return count
}

func (dao *CategoryDaoImpl) GetAllCategories() []db.CategoryDetail {
	//sql := "select category.id, category.category_name,count(article.id) as articleCount from category_tab category left join (select * from article_tab where is_delete = 0 and status in (1,2)) article on category.id = article.category_id group by category.id"
	var categories []db.CategoryDetail
	dao.db.Table("category_tab category").Select("category.id as category_id, category.category_name, count(article.id) as article_count").Joins("left join (select * from article_tab where is_delete = 0 and status in (1,2)) article on category.id = article.category_id group by category.id").Scan(&categories)

	return categories
}

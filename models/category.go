package models

type Category struct {
	Id           int64  `gorm:"id"`
	CategoryName string `gorm:"category_name"` // 分类名
	CreateTime   uint32 `gorm:"create_time"`   // 创建时间
	UpdateTime   uint32 `gorm:"update_time"`   // 更新时间
}

func (Category) TableName() string {
	return "category_tab"
}

type CategoryDetail struct {
	CategoryId   int64  `gorm:"category_id"`
	CategoryName string `gorm:"category_name"`
	ArticleCount int64  `gorm:"article_count"`
}

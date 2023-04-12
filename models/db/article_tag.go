package db

type ArticleTag struct {
	Id        int64 `gorm:"id"`
	ArticleId int64 `gorm:"article_id"` // 文章id
	TagId     int64 `gorm:"tag_id"`     // 标签id
}

func (ArticleTag) TableName() string {
	return "article_tag_tab"
}

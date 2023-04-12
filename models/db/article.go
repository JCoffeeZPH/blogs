package db

type Article struct {
	Id             int64  `gorm:"id"`
	UserId         int64  `gorm:"user_id"`         // 作者
	CategoryId     int64  `gorm:"category_id"`     // 文章分类
	ArticleCover   string `gorm:"article_cover"`   // 文章缩略图
	ArticleTitle   string `gorm:"article_title"`   // 标题
	ArticleContent string `gorm:"article_content"` // 内容
	IsTop          int8   `gorm:"is_top"`          // 是否置顶 0否 1是
	IsFeatured     int8   `gorm:"is_featured"`     // 是否推荐 0否 1是
	IsDelete       int8   `gorm:"is_delete"`       // 是否删除  0否 1是
	Status         int8   `gorm:"status"`          // 状态值 1公开 2私密 3草稿
	Type           int8   `gorm:"type"`            // 文章类型 1原创 2转载 3翻译
	Password       string `gorm:"password"`        // 访问密码
	OriginalUrl    int32  `gorm:"original_url"`    // 原文链接
	CreateTime     uint32 `gorm:"create_time"`     // 发表时间
	UpdateTime     uint32 `gorm:"update_time"`     // 更新时间
}

func (Article) TableName() string {
	return "article_tab"
}

type ArticleWithCategory struct {
	Article
	CategoryName string `gorm:"category_name"`
}

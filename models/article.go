package models

type Article struct {
	Id             int64  `gorm:"id"`
	UserId         int64  `gorm:"user_id"`         // 作者
	CategoryId     int64  `gorm:"category_id"`     // 文章分类
	ArticleCover   string `gorm:"article_cover"`   // 文章缩略图
	ArticleTitle   string `gorm:"article_title"`   // 标题
	ArticleContent string `gorm:"article_content"` // 内容
	IsTop          int64  `gorm:"is_top"`          // 是否置顶 0否 1是
	IsFeatured     int64  `gorm:"is_featured"`     // 是否推荐 0否 1是
	IsDelete       int64  `gorm:"is_delete"`       // 是否删除  0否 1是
	Status         int64  `gorm:"status"`          // 状态值 1公开 2私密 3草稿
	Type           int64  `gorm:"type"`            // 文章类型 1原创 2转载 3翻译
	Password       string `gorm:"password"`        // 访问密码
	OriginalUrl    int32  `gorm:"original_url"`    // 原文链接
	CreateTime     uint32 `gorm:"create_time"`     // 发表时间
	UpdateTime     uint32 `gorm:"update_time"`     // 更新时间
}

func (Article) TableName() string {
	return "article_tab"
}

type ArticleCard struct {
	Id             int64  `gorm:"id"`
	ArticleCover   string `gorm:"articleCover"`
	ArticleTitle   string `gorm:"articleTitle"`
	ArticleContent string `gorm:"articleContent"`
	IsTop          int8   `gorm:"isTop"`
	IsFeatured     int8   `gorm:"isFeatured"`
	Author         struct {
		Id          int64  `gorm:"id"`
		Email       string `gorm:"email"`
		Nickname    string `gorm:"nickname"`
		Avatar      string `gorm:"avatar"`
		Intro       string `gorm:"intro"`
		Website     string `gorm:"website"`
		IsSubscribe int8   `gorm:"isSubscribe"`
		IsDisable   int8   `gorm:"isDisable"`
		CreateTime  uint32 `gorm:"createTime"`
		UpdateTime  uint32 `gorm:"updateTime"`
	}
	CategoryName string `gorm:"categoryName"`
	Tags         []struct {
		Id         int64  `gorm:"id"`
		TagName    string `gorm:"tagName"`
		CreateTime uint32 `gorm:"createTime"`
		UpdateTime uint32 `gorm:"updateTime"`
	}
	Status     int8   `gorm:"status"`
	CreateTime uint32 `gorm:"createTime"`
	UpdateTime uint32 `gorm:"updateTime"`
}

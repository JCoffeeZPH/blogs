package db

type WebsiteConfig struct {
	Id         int64  `gorm:"id"`
	Config     string `gorm:"config"`      // 配置信息
	CreateTime uint32 `gorm:"create_time"` // 创建时间
	UpdateTime uint32 `gorm:"update_time"` // 更新时间
}

func (WebsiteConfig) TableName() string {
	return "website_config_tab"
}

package models

type Tag struct {
	Id         int64  `gorm:"id"`
	TagName    string `gorm:"tag_name"`
	CreateTime uint32 `gorm:"create_time"`
	UpdateTime uint32 `gorm:"update_time"`
}

func (Tag) TableName() string {
	return "tag_tab"
}

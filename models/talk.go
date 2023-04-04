package models

type Talk struct {
	Id         int64  `gorm:"id"`
	UserId     int64  `gorm:"user_id"`
	Content    string `gorm:"content"`
	Images     string `gorm:"images"`
	IsTop      int32  `gorm:"isTop"`
	Status     int32  `gorm:"status"`
	CreateTime uint32 `gorm:"create_time"`
	UpdateTime uint32 `gorm:"update_time"`
}

func (Talk) TableName() string {
	return "talk_tab"
}

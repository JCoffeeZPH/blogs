package db

type UniqueView struct {
	Id         int64  `gorm:"id"`
	ViewCount  int64  `gorm:"view_count"`
	CreateTime uint32 `gorm:"create_time"`
	UpdateTime uint32 `gorm:"update_time"`
}

func (UniqueView) TableName() string {
	return "unique_view_tab"
}

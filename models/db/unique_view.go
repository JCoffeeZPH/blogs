package db

type UniqueView struct {
	Id         int64  `gorm:"id"`
	ViewsCount int64  `gorm:"views_count"`
	Ctime      uint32 `gorm:"ctime"`
}

func (UniqueView) TableName() string {
	return "unique_view_tab"
}

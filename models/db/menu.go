package db

type Menu struct {
	Id         int64  `gorm:"id"`
	Name       string `gorm:"name"`
	Path       string `gorm:"path"`
	Component  string `gorm:"component"`
	Icon       string `gorm:"icon"`
	CreateTime uint32 `gorm:"create_time"`
	UpdateTime uint32 `gorm:"update_time"`
	OrderNum   int32  `gorm:"order_num"`
	ParentId   int64  `gorm:"parent_id"`
	IsHidden   int8   `gorm:"is_hidden"`
}

func (Menu) TableName() string {
	return "menu_tab"
}

package db

type UserInfo struct {
	Id          int64  `gorm:"id"`
	Email       string `gorm:"email"`
	Nickname    string `gorm:"nickname"`
	Avatar      string `gorm:"avatar"`
	Intro       string `gorm:"intro"`
	Website     string `gorm:"website"`
	IsSubscribe int8   `gorm:"is_subscribe"`
	IsDisable   int8   `gorm:"is_disable"`
	CreateTime  uint32 `gorm:"create_time"`
	UpdateTime  uint32 `gorm:"update_time"`
}

func (UserInfo) TableName() string {
	return "user_info_tab"
}

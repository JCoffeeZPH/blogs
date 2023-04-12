package db

type UserAuth struct {
	Id            int64  `gorm:"id"`
	UserInfoId    int64  `gorm:"user_info_id"`
	Username      string `gorm:"username"`
	Password      string `gorm:"password"`
	LoginType     int8   `gorm:"login_type"`
	IpAddress     string `gorm:"ip_address"`
	IpSource      string `gorm:"ip_source"`
	CreateTime    uint32 `gorm:"create_time"`
	UpdateTime    uint32 `gorm:"update_time"`
	LastLoginTime uint32 `gorm:"last_login_time"`
}

func (UserAuth) TableName() string {
	return "user_auth_tab"
}

type UserDetails struct {
	Username      string `gorm:"username"`
	UserInfoId    int64  `gorm:"user_info_id"`
	LoginType     int8   `gorm:"login_type"`
	IpAddress     string `gorm:"ip_address"`
	IpSource      string `gorm:"ip_source"`
	LastLoginTime uint32 `gorm:"last_login_time"`
	Email         string `gorm:"email"`
	Nickname      string `gorm:"nickname"`
	Avatar        string `gorm:"avatar"`
	Intro         string `gorm:"intro"`
	Website       string `gorm:"website"`
	IsSubscribe   int8   `gorm:"is_subscribe"`
	IsDisable     int8   `gorm:"is_disable"`
	RoleId        int8   `gorm:"role_id"`
}

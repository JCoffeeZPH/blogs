package models

type PhotoAlbum struct {
	Id         int64  `gorm:"id"`          // 主键
	AlbumName  string `gorm:"album_name"`  // 相册名
	AlbumDesc  string `gorm:"album_desc"`  // 相册描述
	AlbumCover string `gorm:"album_cover"` // 相册封面
	IsDelete   int64  `gorm:"is_delete"`   // 是否删除
	Status     int64  `gorm:"status"`      // 状态值 1公开 2私密
	CreateTime uint32 `gorm:"create_time"` // 创建时间
	UpdateTime uint32 `gorm:"update_time"` // 更新时间
}

func (PhotoAlbum) TableName() string {
	return "photo_album_tab"
}

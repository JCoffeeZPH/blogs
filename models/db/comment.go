package db

type Comment struct {
	Id             int64  `gorm:"id"`
	UserId         int64  `gorm:"user_id"`
	TopicId        int64  `gorm:"topic_id"`
	CommentContent string `gorm:"comment_content"`
	ReplyUserId    int64  `gorm:"reply_user_id"`
	ParentId       int64  `gorm:"parent_id"`
	Type           int8   `gorm:"type"`
	IsDelete       int8   `gorm:"is_delete"`
	IsReview       int8   `gorm:"is_review"`
	CreateTime     uint32 `gorm:"create_time"`
	UpdateTime     uint32 `gorm:"update_time"`
}

func (Comment) TableName() string {
	return "comment_tab"
}

type CountCommentParam struct {
	Type     *int8
	IsReview *int8
	Nickname *string
}

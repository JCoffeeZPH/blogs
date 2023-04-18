package constants

import "time"

const (
	DefaultWebsiteConfigKey = "blog.website.config.2"
	BlogsViewCountKey       = "blog.view.count"

	UniqueVisitorKey = "blog.unique.visitors"

	// VisitorsAreaDetailKey 用户访问blog、测试账号等登入admin端时保存
	VisitorsAreaDetailKey = "blog.visitors.area"
	// UserAreaDetailKey 用户在blog页面注册的数量
	UserAreaDetailKey = "blog.users.area"

	SessionDefaultTimeOut = 7 * 24 * time.Hour
)

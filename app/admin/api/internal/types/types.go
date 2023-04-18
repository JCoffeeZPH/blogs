// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginResponse struct {
	Id            int64  `json:"id"`
	Avatar        string `json:"avatar"`
	Email         string `json:"email"`
	Intro         string `json:"intro"`
	IpAddress     string `json:"ipAddress"`
	IpSource      string `json:"ipSource"`
	IsSubscribe   int8   `json:"isSubscribe"`
	LastLoginTime string `json:"lastLoginTime"`
	LoginType     int8   `json:"loginType"`
	Nickname      string `json:"nickname"`
	Token         string `json:"token"`
	Username      string `json:"username"`
	Website       string `json:"website"`
}

type GetMenusResponse struct {
	Name      string             `json:"name"`
	Path      string             `json:"path"`
	Component string             `json:"component"`
	Icon      string             `json:"icon"`
	Hidden    bool               `json:"hidden"`
	OrderNum  int32              `json:"orderNum"`
	Children  []GetMenusResponse `json:"children"`
}

type GetAreaDataRequest struct {
	Type int64 `form:"type"`
}

type GetAreaDataResponse struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type GetSysInfoResponse struct {
	ViewsCount   int64         `json:"viewsCount"`
	MessageCount int64         `json:"messageCount"`
	UserCount    int64         `json:"userCount"`
	ArticleCount int64         `json:"articleCount"`
	Categories   []Category    `json:"categories"`
	Tags         []Tag         `json:"tags"`
	ArticleStats []ArticleStat `json:"articleStats"`
	UniqueViews  []UniqueView  `json:"uniqueViews"`
	ArticleRanks []ArticleRank `json:"articleRanks"`
}

type Category struct {
	Id           int64  `json:"id"`
	CategoryName string `json:"categoryName"`
	ArticleCount int64  `json:"articleCount"`
}

type Tag struct {
	Id      int64  `json:"id"`
	TagName string `json:"tagName"`
}

type ArticleStat struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

type UniqueView struct {
	Day        string `json:"day"`
	ViewsCount int64  `json:"viewsCount"`
}

type ArticleRank struct {
	ArticleTitle string `json:"articleTitle"`
	ViewsCount   int64  `json:"viewsCount"`
}

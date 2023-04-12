package svc

import (
	commonConfig "blogs/common/config"
	"blogs/dao"
	"blogs/lib/cache"
)

type ServiceContext struct {
	Config commonConfig.Config
	//BlogClient blogclient.Blog
	PhotoAlbumDao    dao.PhotoAlbumDao
	ArticleDao       dao.ArticleDao
	CategoryDao      dao.CategoryDao
	TagDao           dao.TagDao
	WebsiteConfigDao dao.WebsiteConfigDao
	TalkDao          dao.TalkDao
	UserInfoDao      dao.UserInfoDao
	ArticleTagDao    dao.ArticleTagDao
}

func NewServiceContext(c commonConfig.Config, nc *commonConfig.NacosServerConfig) *ServiceContext {
	dao.InitGorm(c.Mysql.UserName, c.Mysql.Password, c.Mysql.Host, c.Mysql.DatabaseName, c.Mysql.Port)
	cache.InitRedis(c.Redis.Host, c.Redis.Password, c.Redis.Port, c.Redis.DB, c.Redis.PoolSize, c.Redis.MinIdleConns, c.Redis.MaxRetries)

	return &ServiceContext{
		Config: c,
		//BlogClient: blogclient.NewBlog(rpc.NewZrpcClient(constants.BlogRPCServerName, c.Name, nc))
		PhotoAlbumDao:    dao.NewPhotoAlbumDao(),
		ArticleDao:       dao.NewArticleDao(),
		CategoryDao:      dao.NewCategoryDao(),
		TagDao:           dao.NewTagDao(),
		WebsiteConfigDao: dao.NewWebsiteConfigDao(),
		TalkDao:          dao.NewTalkDao(),
		UserInfoDao:      dao.NewUserInfoDaoDao(),
		ArticleTagDao:    dao.NewArticleTagDao(),
	}
}

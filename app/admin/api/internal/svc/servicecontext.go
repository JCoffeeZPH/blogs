package svc

import (
	"blogs/app/admin/api/internal/middleware"
	commonConfig "blogs/common/config"
	"blogs/dao"
	"blogs/lib/cache"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config commonConfig.Config

	AuthToken rest.Middleware

	UserAuthDao dao.UserAuthDao
	UserInfoDao dao.UserInfoDao
	MenuDao     dao.MenuDao
	CommentDao  dao.CommentDao
	ArticleDao  dao.ArticleDao
	UniqueDao   dao.UniqueViewDao
	TagDao      dao.TagDao
	CategoryDao dao.CategoryDao
}

func NewServiceContext(c commonConfig.Config, nc *commonConfig.NacosServerConfig) *ServiceContext {
	dao.InitGorm(c.Mysql.UserName, c.Mysql.Password, c.Mysql.Host, c.Mysql.DatabaseName, c.Mysql.Port)
	cache.InitRedis(c.Redis.Host, c.Redis.Password, c.Redis.Port, c.Redis.DB, c.Redis.PoolSize, c.Redis.MinIdleConns, c.Redis.MaxRetries)

	return &ServiceContext{
		Config: c,

		AuthToken: middleware.NewAuthTokenMiddleware().Handle,

		UserAuthDao: dao.NewUserAuthDao(),
		UserInfoDao: dao.NewUserInfoDaoDao(),
		MenuDao:     dao.NewMenuDao(),
		CommentDao:  dao.NewCommentDao(),
		ArticleDao:  dao.NewArticleDao(),
		UniqueDao:   dao.NewUniqueViewDao(),
		TagDao:      dao.NewTagDao(),
		CategoryDao: dao.NewCategoryDao(),
	}
}

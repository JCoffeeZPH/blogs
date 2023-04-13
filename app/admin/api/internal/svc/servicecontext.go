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
	MenuDao     dao.MenuDao
}

func NewServiceContext(c commonConfig.Config, nc *commonConfig.NacosServerConfig) *ServiceContext {
	dao.InitGorm(c.Mysql.UserName, c.Mysql.Password, c.Mysql.Host, c.Mysql.DatabaseName, c.Mysql.Port)
	cache.InitRedis(c.Redis.Host, c.Redis.Password, c.Redis.Port, c.Redis.DB, c.Redis.PoolSize, c.Redis.MinIdleConns, c.Redis.MaxRetries)

	return &ServiceContext{
		Config: c,

		AuthToken: middleware.NewAuthTokenMiddleware().Handle,

		UserAuthDao: dao.NewUserAuthDao(),
		MenuDao:     dao.NewMenuDao(),
	}
}

package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var dbClient *gorm.DB

func InitGorm(username, password, host, dbName string, port uint64) {
	client, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{Logger: logger.New(
		//将标准输出作为Writer
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			//设定慢查询时间阈值为1ms
			SlowThreshold: 1 * time.Microsecond,
			//设置日志级别，只有Warn和Info级别会输出慢查询日志
			LogLevel: logger.Warn,
		},
	)})

	if err != nil {
		panic(fmt.Sprintf("init gorm error: %+v", err))
	}

	sqlDB, err := client.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	dbClient = client
}

func GetDBClient() *gorm.DB {
	return dbClient
}

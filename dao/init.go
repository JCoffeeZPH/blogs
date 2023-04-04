package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dbClient *gorm.DB

func InitGorm(username, password, host, dbName string, port uint64) {
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)

	client, err := gorm.Open("mysql", datasource)
	if err != nil {
		panic(fmt.Sprintf("init gorm error: %+v", err))
	}
	client.SingularTable(true)
	sqlDB := client.DB()
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	client.LogMode(true)
	dbClient = client
}

func GetDBClient() *gorm.DB {
	return dbClient
}

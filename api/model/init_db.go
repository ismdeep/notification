package model

import (
	"github.com/ismdeep/notification/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // load mysql driver
	"time"
)

// DB 数据库链接单例
var DB *gorm.DB

// InitDatabase 在中间件中初始化mysql链接
func InitDatabase() {
	db, err := gorm.Open("mysql", config.MySQL.DSN)
	if err != nil {
		panic(err)
	}

	db.LogMode(config.MySQL.LogMode)

	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
}

// AutoMigrate 自动迁移
func AutoMigrate() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Token{})
	DB.AutoMigrate(&Mail{})
}

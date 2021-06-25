package main

import (
	"github.com/ismdeep/notification/api"
	"github.com/ismdeep/notification/load"
)

func main() {
	load.Config()        // 加载配置
	load.DB()            // 加载数据库
	load.DBAutoMigrate() // 数据库自动迁移
	load.MailService()   // 邮件服务初始化

	api.LoadAPIServer(false) // 加载API服务
}

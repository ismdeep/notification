package main

import (
	"github.com/ismdeep/notification/api/controller"
	"github.com/ismdeep/notification/load"
)

func main() {
	load.Config()                   // 加载配置
	load.JWT()                      // 加载JWT
	load.DB()                       // 加载数据库
	load.DBAutoMigrate()            // 数据库自动迁移
	load.MailService()              // 邮件服务初始化
	controller.LoadAPIServer(false) // 加载API服务
}

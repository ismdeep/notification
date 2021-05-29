package main

import (
	"github.com/ismdeep/args"
	"github.com/ismdeep/notification/api"
	"github.com/ismdeep/notification/config"
	"github.com/ismdeep/notification/service/mail"
)

func main() {
	// 加载配置
	config.LoadConfig(args.GetValue("-c"))

	// 邮件服务初始化
	mail.Init()

	// 加载API服务
	api.LoadAPIServer(false)
}

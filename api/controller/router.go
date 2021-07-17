package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/notification/config"
)

// LoadAPIServer 加载API服务器
func LoadAPIServer(detached bool) {
	router := gin.Default()

	router.POST("/api/v1/sign-up", SignUp)         // 注册账号
	router.POST("/api/v1/sign-in", SignIn)         // 登录
	router.GET("/api/v1/user/info", GetUserInfo)   // 用户信息
	router.GET("/api/v1/mail-types", GetMailTypes) // 邮件类型
	router.POST("/api/v1/mails", SendEmail)        // 发送邮件
	router.POST("/api/v1/tokens", GenerateToken)   // 生成Token

	if err := router.Run(config.Global.Bind); err != nil {
		panic(err)
	}
}

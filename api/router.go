package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/notification/api/controller"
	"github.com/ismdeep/notification/config"
)

// LoadAPIServer 加载API服务器
func LoadAPIServer(detached bool) {
	router := gin.Default()
	router.GET("/api/v1/mail-types", controller.GetMailTypes)
	router.POST("/api/v1/mails", controller.PushMail)
	//router.POST("/api/v1/wecom-messages", )
	router.POST("/api/v1/sign-up", controller.SignUp)
	router.POST("/api/v1/sign-in", controller.SignIn)

	if err := router.Run(config.Global.Bind); err != nil {
		panic(err)
	}
}

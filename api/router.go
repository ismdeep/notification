package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/notification/api/controller"
	"github.com/ismdeep/notification/config"
)

func daemon(router *gin.Engine) {
	if err := router.Run(config.Global.Bind); err != nil {
		panic(err)
	}
}

// LoadAPIServer 加载API服务器
func LoadAPIServer(detached bool) {
	router := gin.Default()
	router.GET("/api/v1/mail-types", controller.GetMailTypes)
	router.POST("/api/v1/mails", controller.PushMail)

	if detached {
		go daemon(router)
	} else {
		daemon(router)
	}
}

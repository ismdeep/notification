package api

import (
	"github.com/gin-gonic/gin"
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
	router.GET("/mail-types", GetMailTypes)
	router.POST("/mails", PushMail)

	if detached {
		go daemon(router)
	} else {
		daemon(router)
	}
}

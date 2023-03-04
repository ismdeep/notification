package api

import (
	"github.com/gin-gonic/gin"

	"github.com/ismdeep/notification/app/server/auth"
	"github.com/ismdeep/notification/app/server/handler"
	"github.com/ismdeep/notification/app/server/request"
	"github.com/ismdeep/notification/pkg/core"
)

// SignIn 登录
// @Summary 登录
// @Author l.jiang.1024@gmail.com
// @Description 登录
// @Tags Auth
// @Router	/api/v1/sign-in [put]
func SignIn(c *gin.Context) any {
	var req request.Login
	core.PanicIf(
		c.ShouldBindJSON(&req))
	return handler.Auth.Login(&req)
}

// GetUserInfo 用户信息
// @Summary 用户信息
// @Author l.jiang.1024@gmail.com
// @Description 用户信息
// @Tags Auth
// @Success 200
// @Router	/api/v1/profile [get]
func GetUserInfo(c *gin.Context) any {
	return auth.GetUserInfo(c)
}

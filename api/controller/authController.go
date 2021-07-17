package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/notification/api/auth"
	authHandler "github.com/ismdeep/notification/api/handler/auth"
	"github.com/ismdeep/notification/api/request"
	"github.com/ismdeep/notification/common"
)

// SignUp 注册
// @Summary 注册
// @Author @uniontech.com
// @Description 注册
// @Tags Auth
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Router	/api/v1/sign-up [post]
func SignUp(c *gin.Context) {
	req := &request.Register{}
	if err := c.Bind(req); err != nil {
		JSON(c, WithError(common.ErrBadRequest))
		return
	}

	respData, err := authHandler.Register(req)
	if err != nil {
		JSON(c, WithError(err))
		return
	}

	JSON(c, WithData(respData))
	return
}

// SignIn 登录
// @Summary 登录
// @Author @uniontech.com
// @Description 登录
// @Tags Auth
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Router	/api/v1/sign-in [put]
func SignIn(c *gin.Context) {
	req := &request.Login{}
	if err := c.Bind(req); err != nil {
		JSON(c, WithError(err))
		return
	}

	respData, err := authHandler.Login(req)
	if err != nil {
		JSON(c, WithError(err))
		return
	}

	JSON(c, WithData(respData))
	return
}

// GetUserInfo 用户信息
// @Summary 用户信息
// @Author jianglinwei@uniontech.com
// @Description 用户信息
// @Tags Auth
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body	request. true "JSON数据"
// @Success 200 {object} response.
// @Router	/api/v1/user/info [get]
func GetUserInfo(c *gin.Context) {
	userInfo, err := auth.GetUserInfo(c)
	if err != nil {
		JSON(c, WithError(err))
		return
	}

	JSON(c, WithData(userInfo))
	return
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/notification/api/handler/auth"
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

	respData, err := auth.Register(req)
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

	respData, err := auth.Login(req)
	if err != nil {
		JSON(c, WithError(err))
		return
	}

	JSON(c, WithData(respData))
	return
}

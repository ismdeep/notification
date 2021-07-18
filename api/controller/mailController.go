package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/notification/api/auth"
	mailHandler "github.com/ismdeep/notification/api/handler/mail"
	"github.com/ismdeep/notification/api/request"
)

// GetMailTypes 获取邮件类型
// @Summary 获取邮件类型
// @Author l.jiang.1024@gmail.com
// @Description 获取邮件类型
// @Tags mail
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Router	/api/v1/mail-types [get]
func GetMailTypes(c *gin.Context) {
	respData := mailHandler.GetMailTypes()
	c.JSON(0, respData)
	return
}

// SendEmail 发送邮件
// @Summary 发送邮件
// @Author l.jiang.1024@gmail.com
// @Description 发送邮件
// @Tags Mail
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body	request.PushMailRequest true "JSON数据"
// @Router	/api/v1/mails [post]
func SendEmail(c *gin.Context) {

	userInfo, err := auth.GetUserInfo(c)
	if err != nil {
		JSON(c, WithError(err))
		return
	}

	req := &request.PushMailRequest{}
	if err := c.Bind(req); err != nil {
		JSON(c, WithError(errors.New("bind request failed")))
		return
	}

	if err := mailHandler.AddMail(userInfo.ID, req); err != nil {
		JSON(c, WithError(err))
		return
	}

	JSON(c, WithMsg("mail sent"))
	return
}

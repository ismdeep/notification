package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/notification/api/auth"
	"github.com/ismdeep/notification/api/handler/token"
	"github.com/ismdeep/notification/api/request"
	"github.com/ismdeep/notification/common"
)

// GenerateToken 生成Token
// @Summary 生成Token
// @Author jianglinwei@uniontech.com
// @Description 生成Token
// @Tags Token
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param body body	request.NewToken true "JSON数据"
// @Success 200 {object} response.TokenDetail
// @Router	/api/v1/tokens [post]
func GenerateToken(c *gin.Context) {
	req := &request.NewToken{}
	_ = c.Bind(req)

	userInfo, err := auth.GetUserInfo(c)
	if err != nil {
		JSON(c, WithError(common.ErrNotLogin))
		return
	}

	newToken, err := token.NewToken(userInfo.ID, req)
	if err != nil {
		JSON(c, WithError(err))
		return
	}

	JSON(c, WithData(newToken))
	return
}

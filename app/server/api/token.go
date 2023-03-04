package api

import (
	"github.com/gin-gonic/gin"

	"github.com/ismdeep/notification/app/server/auth"
	"github.com/ismdeep/notification/app/server/handler"
	"github.com/ismdeep/notification/app/server/request"
	"github.com/ismdeep/notification/app/server/store"
	"github.com/ismdeep/notification/pkg/core"
)

// GenerateToken 生成Token
// @Summary 生成Token
// @Author l.jiang.1024@gmail.com
// @Description 生成Token
// @Tags Token
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param group_id path integer          true "分组ID"
// @Param req body	request.NewToken true "JSON数据"
// @Success 200
// @Router /api/v1/groups/:group_id/tokens [post]
func GenerateToken(c *gin.Context) any {
	var req request.NewToken
	core.PanicIf(c.ShouldBindJSON(&req))
	return handler.Token.Create(auth.GetUserInfo(c).ID, req)
}

// GetTokenList 获取Token列表
// @Summary 获取Token列表
// @Author l.jiang.1024@gmail.com
// @Description 获取Token列表
// @Tags Token
// @Param Authorization	header	string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param group_id path integer          true "分组ID"
// @Param req      body request.NewToken true "JSON数据"
// @Success 200
// @Router /api/v1/tokens [post]
func GetTokenList(c *gin.Context) any {
	return store.Token.ListByUserID(auth.GetUserInfo(c).ID)
}

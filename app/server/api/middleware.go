package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ismdeep/notification/app/server/auth"
)

// Auth 用户访问检查
func Auth(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": r.(error).Error(),
			})
			c.Abort()
			return
		}
	}()
	auth.GetUserInfo(c)
	c.Next()
}

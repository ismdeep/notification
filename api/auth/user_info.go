package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// UserInfo user info
func UserInfo(c *gin.Context) {
	authKey := c.GetHeader("Authorization")
	fmt.Println(authKey)
}

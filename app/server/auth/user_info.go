package auth

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ismdeep/notification/pkg/core"
	"github.com/ismdeep/notification/pkg/jwt"
)

// GroupInfo 邮件群组信息
type GroupInfo struct {
	ID uint `json:"id"`
}

// UserInfo 用户信息
type UserInfo struct {
	ID       string `json:"id"`       // ID
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) *UserInfo {
	// 1. 从 Authorization 中获取用户信息
	jwtToken := c.GetHeader("Authorization")
	if strings.Contains(jwtToken, "Bearer ") {
		jwtToken = strings.Split(jwtToken, " ")[1]
		if userInfo, err := JWTVerify(jwtToken); err == nil {
			return userInfo
		}
		core.PanicIf(errors.New("尚未登录"))
	}

	core.PanicIf(errors.New("尚未登录"))
	return nil
}

// JWTSign jwt签名
func JWTSign(userInfo *UserInfo) (string, error) {
	if userInfo == nil {
		return "", errors.New("bad request")
	}

	userInfoBytes, err := json.Marshal(userInfo)
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(string(userInfoBytes))
	if err != nil {
		return "", err
	}

	return token, nil
}

// JWTVerify jwt验证
func JWTVerify(jwtToken string) (*UserInfo, error) {
	token, err := jwt.VerifyToken(jwtToken)
	if err != nil {
		return nil, err
	}

	var userInfo UserInfo
	if err := json.Unmarshal([]byte(token), &userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}

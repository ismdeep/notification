package auth

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/jwt"
	"github.com/ismdeep/notification/api/model"
	"github.com/ismdeep/notification/common"
	"github.com/jinzhu/gorm"
	"strings"
)

// UserInfo 用户信息
type UserInfo struct {
	ID       uint   `json:"id"`       // ID
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) (*UserInfo, error) {

	// 1. 从 Authorization 中获取用户信息
	jwtToken := c.GetHeader("Authorization")
	if strings.Contains(jwtToken, "Bearer ") {
		jwtToken = strings.Split(jwtToken, " ")[1]
		if userInfo, err := JWTVerify(jwtToken); err == nil {
			return userInfo, nil
		}
		return nil, common.ErrNotLogin
	}

	// 2. 从 Token 中获取用户信息
	tokenKey := c.GetHeader("Token")
	if tokenKey != "" {
		if userInfo, err := TokenVerify(tokenKey); err == nil {
			return userInfo, nil
		}
		return nil, common.ErrNotLogin
	}

	return nil, common.ErrNotLogin
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

	userInfo := &UserInfo{}
	if err := json.Unmarshal([]byte(token), userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}

// TokenVerify token验证
func TokenVerify(tokenStr string) (*UserInfo, error) {
	token := &model.Token{}

	if err := model.DB.Where("token = ?", tokenStr).First(token).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("token error")
		}
		return nil, common.ErrSystemError
	}

	user := &model.User{}

	if err := model.DB.Where("id = ?", token.UserID).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrUserNotExists
		}

		return nil, common.ErrSystemError
	}

	return &UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}, nil
}

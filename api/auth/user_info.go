package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
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
	authKey := c.GetHeader("Authorization")
	if strings.Contains(authKey, "Bearer ") {
		return nil, common.ErrNotImplemented
	}

	tokenKey := c.GetHeader("Token")
	if tokenKey == "" {
		return nil, common.ErrNotLogin
	}
	userInfo, err := TokenVerify(tokenKey)
	if err != nil {
		return nil, common.ErrNotLogin
	}

	return userInfo, nil
}

// JWTSign jwt签名
func JWTSign(userInfo *UserInfo) (string, error) {
	if userInfo == nil {
		return "", errors.New("bad request")
	}

	//marshal, err := json.Marshal(userInfo)
	//if err != nil {
	//	return "", err
	//}

	return "", common.ErrNotImplemented
}

// JWTVerify jwt验证
func JWTVerify(jwtToken string) (*UserInfo, error) {
	return nil, common.ErrNotImplemented
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

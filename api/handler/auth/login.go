package auth

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ismdeep/notification/api/model"
	"github.com/ismdeep/notification/api/request"
	"github.com/ismdeep/notification/api/response"
	"github.com/ismdeep/notification/common"
	"github.com/ismdeep/notification/config"
	"github.com/jinzhu/gorm"
	"time"
)

// JWTUserInfo jwt user info
type JWTUserInfo struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}

// Login 登录
func Login(req *request.Login) (*response.Login, error) {
	if req == nil {
		return nil, common.ErrBadRequest
	}

	if req.Username == "" {
		return nil, common.ErrBadRequest
	}

	// 1. 检查用户是否存在
	user := &model.User{}
	if err := model.DB.Where("username = ?", req.Username).First(user).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, common.ErrSystemError
		}

		return nil, common.ErrUserNotExists
	}

	// 2. 检查密码
	auth := &model.Auth{}
	if err := model.DB.Where("user_id = ?", user.ID).First(auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("请设置密码")

		}

		return nil, common.ErrSystemError
	}

	if auth.Digest != fmt.Sprintf("%x", md5.Sum([]byte(req.Password))) {
		return nil, errors.New("密码错误")
	}

	// 3. 检查用户是否被禁用
	if !user.Enabled {
		return nil, common.ErrUserHasBeenBaned
	}

	// 4. 签名
	binData, err := json.Marshal(&JWTUserInfo{
		UserID:   user.ID,
		Username: user.Username,
	})
	if err != nil {
		return nil, common.ErrSystemError
	}
	expireDuration, err := time.ParseDuration(config.JWT.Expire)
	if err != nil {
		return nil, common.ErrSystemError
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": string(binData),
		"exp": time.Now().Add(expireDuration).Unix(),
	})
	accessToken, err := token.SignedString([]byte(config.JWT.Key))
	if err != nil {
		return nil, common.ErrSystemError
	}

	return &response.Login{
		UserID:      user.ID,
		AccessToken: accessToken,
	}, nil
}

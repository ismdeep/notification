package auth

import (
	"errors"
	"github.com/ismdeep/digest"
	"github.com/ismdeep/notification/api/auth"
	"github.com/ismdeep/notification/api/model"
	"github.com/ismdeep/notification/api/request"
	"github.com/ismdeep/notification/api/response"
	"github.com/ismdeep/notification/common"
	"github.com/jinzhu/gorm"
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
	if !digest.Verify(user.Digest, req.Password) {
		return nil, errors.New("密码错误")
	}

	// 3. 检查用户是否被禁用
	if !user.Enabled {
		return nil, common.ErrUserHasBeenBaned
	}

	// 4. 签名
	jwtToken, err := auth.JWTSign(&auth.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	})
	if err != nil {
		return nil, err
	}

	return &response.Login{
		UserID:      user.ID,
		AccessToken: jwtToken,
	}, nil
}

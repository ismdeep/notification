package auth

import (
	"crypto/md5"
	"fmt"
	"github.com/ismdeep/notification/api/request"
	"github.com/ismdeep/notification/api/response"
	"github.com/ismdeep/notification/common"
	"github.com/ismdeep/notification/model"
)

// Register 注册
func Register(req *request.Register) (*response.Register, error) {
	if req == nil {
		return nil, common.ErrBadRequest
	}

	if err := req.Check(); err != nil {
		return nil, common.ErrBadRequest
	}

	if req.Nickname == "" {
		req.Nickname = req.Username
	}

	user := &model.User{
		Username: req.Username,
		Nickname: req.Nickname,
		Avatar:   "",
		Enabled:  true,
	}

	// 写入用户
	if err := model.DB.Save(user).Error; err != nil {
		return nil, common.ErrSystemError
	}

	// 写入密码
	auth := &model.Auth{
		UserID:  user.ID,
		Digest:  fmt.Sprintf("%x", md5.Sum([]byte(req.Password))),
	}
	if err := model.DB.Save(auth).Error; err != nil {
		return nil, common.ErrSystemError
	}

	return &response.Register{UserID: user.ID}, nil
}

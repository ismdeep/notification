package auth

import (
	"github.com/ismdeep/digest"
	"github.com/ismdeep/notification/api/model"
	"github.com/ismdeep/notification/api/request"
	"github.com/ismdeep/notification/api/response"
	"github.com/ismdeep/notification/common"
)

// Register 注册
func Register(req *request.Register) (*response.Register, error) {
	if req == nil {
		return nil, common.ErrBadRequest
	}

	if err := req.Check(); err != nil {
		return nil, common.ErrBadRequest
	}

	user := &model.User{
		Username: req.Username,
		Nickname: req.Username,
		Avatar:   "",
		Enabled:  true,
		Digest:   digest.Generate(req.Password),
	}

	// 写入用户
	if err := model.DB.Create(user).Error; err != nil {
		return nil, common.ErrDatabaseOperateFailed
	}

	return &response.Register{UserID: user.ID}, nil
}

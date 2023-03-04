package handler

import (
	"errors"

	"github.com/ismdeep/digest"

	"github.com/ismdeep/notification/app/server/auth"
	"github.com/ismdeep/notification/app/server/request"
	"github.com/ismdeep/notification/app/server/response"
	"github.com/ismdeep/notification/app/server/store"
	"github.com/ismdeep/notification/pkg/core"
)

type authHandler struct {
}

// Auth handler
var Auth *authHandler

// JWTUserInfo jwt user info
type JWTUserInfo struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}

// Login 登录
func (receiver *authHandler) Login(req *request.Login) *response.Login {
	// 0. 参数检查
	core.PanicIf(
		core.IfErr(req == nil, errors.New("bad request")))
	core.PanicIf(
		core.IfErr(req.Username == "", errors.New("bad request, parameter is empty: username")))

	// 1. 检查用户是否存在
	user := store.User.GetByUsername(req.Username)

	// 2. 检查密码
	core.PanicIf(
		core.IfErr(!digest.Verify(user.Digest, req.Password), errors.New("认证失败")))

	// 3. 检查用户是否被禁用
	core.PanicIf(
		core.IfErr(!user.Enabled, errors.New("用户已被封禁")))

	// 4. 签名
	jwtToken, err := auth.JWTSign(&auth.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
	})
	core.PanicIf(err)

	return &response.Login{
		UserID:      user.ID,
		AccessToken: jwtToken,
	}
}

// Register 注册
func (receiver *authHandler) Register(req request.Register) *response.Register {
	core.PanicIf(
		req.Check())

	userID := store.User.Create(req.Username, req.Password)
	return &response.Register{UserID: userID}
}

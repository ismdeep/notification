package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/notification/config"
)

// UserInfo user info
type UserInfo struct {
	ID       uint
	Username string
	Nickname string
}

// AuthUser auth user
func AuthUser(c *gin.Context) *UserInfo {
	authKey := c.GetHeader("Authorization")
	if authKey != fmt.Sprintf("Bearer %v", config.Global.Secret) {
		return nil
	}

	// @TODO
	return &UserInfo{
		ID:       0,
		Username: "",
		Nickname: "",
	}
}

// Option option
type Option struct {
	Code *int
	Msg  *string
	Data *interface{}
}

// WithMsg With Msg
func WithMsg(msg string) *Option {
	return &Option{
		Msg: &msg,
	}
}

// WithError With error
func WithError(err error) *Option {
	val := 1
	msg := err.Error()
	return &Option{
		Code: &val,
		Msg:  &msg,
	}
}

// WithData With Data
func WithData(data interface{}) *Option {
	return &Option{
		Data: &data,
	}
}

func renderRespData(defaultCode int, options ...*Option) map[string]interface{} {
	respData := make(map[string]interface{})
	respData["code"] = defaultCode

	for _, option := range options {
		if option.Code != nil {
			respData["code"] = option.Code
		}
		if option.Msg != nil {
			respData["msg"] = option.Msg
		}
		if option.Data != nil {
			respData["data"] = option.Data
		}
	}

	return respData
}

// JSON JSON
func JSON(c *gin.Context, options ...*Option) {
	respData := renderRespData(0, options...)
	c.JSON(0, respData)
}

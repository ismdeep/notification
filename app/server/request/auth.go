package request

import "errors"

// Register 注册请求
type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Check 注册请求参数检查
func (receiver *Register) Check() error {
	if receiver.Username == "" {
		return errors.New("username can not be empty")
	}

	if len(receiver.Password) < 6 {
		return errors.New("password length must not less than 6")
	}

	return nil
}

// Login 登录请求
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Check 登录请求参数检查
func (receiver *Login) Check() error {
	if receiver.Username == "" {
		return errors.New("username can not be empty")
	}

	return nil
}

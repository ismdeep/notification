package response

// Register 注册返回
type Register struct {
	UserID uint `json:"user_id"`
}

// Login 登陆返回
type Login struct {
	UserID      uint   `json:"user_id"`
	AccessToken string `json:"access_token"`
}

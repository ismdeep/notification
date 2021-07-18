package load

import (
	"github.com/ismdeep/jwt"
	"github.com/ismdeep/notification/config"
)

// JWT 加载JWT
func JWT() {
	err := jwt.Init(config.JWT.Key, config.JWT.Expire)
	if err != nil {
		panic(err)
	}
}

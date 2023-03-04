package conf

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"

	"github.com/ismdeep/notification/internal/solidutil"
	"github.com/ismdeep/notification/pkg/jwt"
)

// s 配置结构体
type s struct {
	Server struct {
		Bind string `env:"SERVER_BIND,default=0.0.0.0:9000"`
	}
	DB struct {
		Dialect string `env:"DB_DIALECT,default=mysql"`
		DSN     string `env:"DB_DSN,default=root:notification123456@tcp(127.0.0.1:3306)/notification?parseTime=true&loc=Local&charset=utf8mb4,utf8"`
	}
	Security struct {
		JWT              string `env:"SECURITY_JWT"`
		TelegramBotToken string `env:"SECURITY_TELEGRAM_BOT_TOKEN"`
		TelegramChatID   string `env:"SECURITY_TELEGRAM_CHAT_ID"`
	}
	Proxy struct {
		Socks5         string `env:"PROXY_SOCKS5"`
		Socks5Username string `env:"PROXY_SOCKS5_USERNAME"`
		Socks5Password string `env:"PROXY_SOCKS5_PASSWORD"`
	}
}

// ROOT instance
var ROOT s

func init() {
	if err := envconfig.Process(context.Background(), &ROOT); err != nil {
		panic(err)
	}

	// check proxy
	if ROOT.Proxy.Socks5 != "" {
		fmt.Println("[INFO] system is using socks5 proxy:", ROOT.Proxy.Socks5)
	}

	// init jwt
	if ROOT.Security.JWT == "" {
		fmt.Println("[WARN] system is using a RANDOM JWT KEY, du env is empty: SECURITY_JWT")
		ROOT.Security.JWT = solidutil.RandToken()
	}
	jwt.Init(&jwt.Config{
		Key:    ROOT.Security.JWT,
		Expire: "72h",
	})
}

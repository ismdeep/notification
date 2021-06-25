package load

import (
	"github.com/ismdeep/args"
	"github.com/ismdeep/notification/config"
	"os"
)

// Config 加载配置
func Config() {
	configPath := ""
	configPath = args.GetValue("-c")
	if configPath == "" {
		configPath = os.Getenv("NOTIFICATION_CONFIG_TOML")
	}

	if configPath == "" {
		panic("config path is not specified")
	}

	config.LoadConfig(configPath)
}

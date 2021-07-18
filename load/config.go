package load

import (
	"github.com/ismdeep/args"
	"github.com/ismdeep/notification/config"
)

// Config 加载配置
func Config() {
	configPath := "./configs/server.toml"

	if args.Exists("-c") {
		configPath = args.GetValue("-c")
	}

	config.LoadConfig(configPath)
}

package load

import (
	"fmt"
	"github.com/ismdeep/notification/config"
	"os"
)

// Config 加载配置
func Config(fileName string) {
	configFolderPath := os.Getenv("NOTIFICATION_CONFIGS")
	if configFolderPath == "" {
		panic("config path is not specified")
	}

	configPath := fmt.Sprintf("%v/%v", configFolderPath, fileName)
	config.LoadConfig(configPath)
}

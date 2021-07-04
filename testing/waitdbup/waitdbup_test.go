package waitdbup

import (
	"github.com/ismdeep/notification/config"
	"github.com/ismdeep/notification/load"
	"github.com/jinzhu/gorm"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	load.Config(config.TestConfig)

	for {
		_, err := gorm.Open("mysql", config.MySQL.DSN)
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		break
	}

	m.Run()
}

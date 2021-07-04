package mail

import (
	"github.com/ismdeep/notification/config"
	"github.com/ismdeep/notification/load"
	"testing"
)

func TestMain(m *testing.M) {
	load.Config(config.TestConfig)
	load.DB()

	m.Run()
}

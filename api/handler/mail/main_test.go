package mail

import (
	"github.com/ismdeep/notification/load"
	"testing"
)

func TestMain(m *testing.M) {
	load.Config()
	load.DB()

	m.Run()
}

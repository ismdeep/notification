package auth

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ismdeep/notification/load"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	load.Config()
	load.DB()

	gofakeit.Seed(time.Now().UnixNano())

	m.Run()
}

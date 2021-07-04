package prepare

import (
	authHandler "github.com/ismdeep/notification/api/handler/auth"
	tokenHandler "github.com/ismdeep/notification/api/handler/token"
	"github.com/ismdeep/notification/api/request"
	"github.com/ismdeep/notification/config"
	"github.com/ismdeep/notification/load"
	"testing"
)

func TestMain(m *testing.M) {
	load.Config(config.TestConfig)
	load.DB()
	load.DBAutoMigrate()

	var err error

	_, err = authHandler.Register(&request.Register{
		Username: "user001",
		Nickname: "L. Jiang",
		Password: "1234567890",
	})
	if err != nil {
		panic(err)
	}

	_, err = tokenHandler.NewToken(1, &request.NewToken{
		Name: "ALL STAR",
	})
	if err != nil {
		panic(err)
	}

	m.Run()
}

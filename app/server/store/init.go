package store

import (
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ismdeep/notification/app/server/conf"
	"github.com/ismdeep/notification/pkg/core"
)

// 数据库连接
var db *gorm.DB

//go:embed migrations/*.sql
var sqlFS embed.FS

// init 在中间件中初始化mysql链接
func init() {
	switch conf.ROOT.DB.Dialect {
	case "mysql":
		conn, err := gorm.Open(mysql.Open(conf.ROOT.DB.DSN))
		core.PanicIf(
			err)
		db = conn
	default:
		panic(fmt.Sprintf("[ERROR] unsupported db dialect: %v", conf.ROOT.DB.Dialect))
	}

	// migration
	goose.SetBaseFS(sqlFS)
	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	core.PanicIf(err)
	if err := goose.Up(sqlDB, "migrations"); err != nil {
		panic(err)
	}

	//// AutoMigration
	//core.PanicIf(
	//	db.AutoMigrate(
	//		&model.User{},
	//		&model.Token{},
	//		&model.Project{},
	//		&model.Msg{},
	//	))

	// 自动创建初始用户
	User.InitUser()

}

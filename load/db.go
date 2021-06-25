package load

import "github.com/ismdeep/notification/model"

// DB 加载数据库
func DB() {
	model.InitDatabase()
}

// DBAutoMigrate 加载数据库自动迁移
func DBAutoMigrate() {
	model.AutoMigrate()
}

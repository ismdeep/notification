package model

import "github.com/jinzhu/gorm"

// User user model
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50); not null;unique;"` // 用户名
	Nickname string `gorm:"type:varchar(255);not null;"`        // 昵称
	Avatar   string `gorm:"type:varchar(255);not null;"`        // 头像URL地址或图片Base64编码
	Enabled  bool   `gorm:"type:tinyint;not null;default: 0;"`  // 是否启用
	Digest   string `gorm:"type:varchar(255);not null;"`        // 数字摘要
}

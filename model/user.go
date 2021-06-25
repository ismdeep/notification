package model

import "github.com/jinzhu/gorm"

// User user model
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50); not null;unique;"`
	Nickname string `gorm:"type:varchar(255);not null;"`
	Avatar   string `gorm:"type:varchar(255);not null;"`
	Enabled  bool   `gorm:"type:tinyint;not null;default: 0;"` // 是否启用
}

package model

import "github.com/jinzhu/gorm"

// Auth auth
type Auth struct {
	gorm.Model
	UserID uint   `gorm:"type:int unsigned;not null;unique;"` // 用户ID
	Digest string `gorm:"type:varchar(255);not null;"`        // 数字摘要
}

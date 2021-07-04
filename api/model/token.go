package model

import "github.com/jinzhu/gorm"

// Token token
type Token struct {
	gorm.Model
	UserID    uint   `gorm:"type:int unsigned;not null;"` // 用户ID
	TokenName string `gorm:"type:varchar(255);not null;"` // Token名称
	Token     string `gorm:"type:varchar(255);not null;"` // Token值
}

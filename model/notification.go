package model

import "github.com/jinzhu/gorm"

// Notification notification
type Notification struct {
	gorm.Model
	UserID   uint   `gorm:"type:int unsigned;not null;"`               // 用户ID
	Status   uint   `gorm:"type:tinyint unsigned;not null;default 0;"` // 状态，0未发送，1已发送，2发送失败
	Type     string `gorm:"type:varchar(50);not null;"`                // 类型，目前支持：email
	PackData string `gorm:"type:text;not null;"`                       // 消息包JSON字符串
}

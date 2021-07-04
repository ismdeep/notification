package model

import "github.com/jinzhu/gorm"

// Mail Mail
type Mail struct {
	gorm.Model
	UserID     uint   `gorm:"type:int unsigned;not null;"`               // 用户ID
	Status     uint   `gorm:"type:tinyint unsigned;not null;default 0;"` // 状态，0未发送，1已发送，2发送失败
	SenderName string `gorm:"type:varchar(255);not null;"`               // 发送者名称
	Type       string `gorm:"type:varchar(50);not null;"`                // 类型，支持：text, html
	Subject    string `gorm:"type:varchar(255);not null;"`               // 主题
	Content    string `gorm:"type:text;not null;"`                       // 内容
	ToMailList string `gorm:"type:text(255);not null;"`                  // 收件人列表
}

const (
	// StatusPending 未发送
	StatusPending = 0
	// StatusSent 发送成功
	StatusSent = 1
	// StatusFailed 发送失败
	StatusFailed = 2
)

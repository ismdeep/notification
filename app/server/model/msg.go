package model

const (
	MsgStatusPending = 0
	MsgStatusSent    = 1
	MsgStatusFailed  = 2
)

type Msg struct {
	ID                string `gorm:"type:varchar(255);not null;primaryKey"`
	CustomerMsgID     string `gorm:"type:varchar(255);not null;uniqueIndex:uq_userid_customermsgid"`
	UserID            string `gorm:"type:varchar(255);not null;uniqueIndex:uq_userid_customermsgid"`
	Content           string `gorm:"type:longtext"`
	Status            int    `gorm:"type:tinyint;not null"`
	Err               string `gorm:"type:longtext"`
	CreatedAtUnixNano int64  `gorm:"type:bigint unsigned;not null;index:idx_msgs_created_at_unix_nano"`
}

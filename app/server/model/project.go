package model

type Project struct {
	ID                string `gorm:"type:varchar(255);not null;primaryKey"`
	Name              string `gorm:"type:varchar(255);not null;unique"`
	CreatedAtUnixNano int64  `gorm:"type:bigint unsigned;not null;"`
}

package model

// Token token
// token.id format: tid_<rand-hex>
//
//	generate by:
//	  $ echo "tid_$(openssl rand --hex 16)"
//
// token value format: nt_<rand-hex>
//
//	generate by:
//	  $ echo "nt_$(openssl rand --hex 16)"
type Token struct {
	ID                string `gorm:"type:varchar(255);primaryKey" json:"id"`
	UserID            string `json:"-"`
	TokenName         string `gorm:"type:varchar(255);not null;" json:"token_name"` // token name
	Token             string `gorm:"type:varchar(255);not null;" json:"token"`      // token value, format: nt_7b95153f72af416eb0a8a7f43ea62f80
	AESKey            string `gorm:"type:varchar(255);not null;" json:"aes_key"`    // openssl aes-256-cbc key
	CreatedAtUnixNano int64  `gorm:"type:bigint unsigned;not null" json:"-"`
}

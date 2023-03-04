package store

import (
	"time"

	"github.com/ismdeep/notification/app/server/model"
	"github.com/ismdeep/notification/internal/solidutil"
	"github.com/ismdeep/notification/pkg/core"
)

type tokenStore struct {
}

// Token store
var Token *tokenStore

func (receiver *tokenStore) InitToken(userID string) {
	firstTokenName := "First Token"
	if !Token.ExistsByID(solidutil.TokenID(userID, firstTokenName)) {
		receiver.Create(userID, firstTokenName)
	}
}

func (receiver *tokenStore) GetByID(tokenID string) *model.Token {
	var t model.Token
	core.PanicIf(db.Where("id = ?", tokenID).First(&t).Error)
	return &t
}

func (receiver *tokenStore) ExistsByID(tokenID string) bool {
	var cnt int64
	core.PanicIf(db.Model(&model.Token{}).Where("id = ?", tokenID).Count(&cnt).Error)
	return cnt > 0
}

func (receiver *tokenStore) GetByToken(tokenStr string) *model.Token {
	var token model.Token
	core.PanicIf(
		db.Where("token = ?", tokenStr).Limit(1).First(&token).Error)
	return &token
}

func (receiver *tokenStore) Create(userID string, tokenName string) *model.Token {
	t := model.Token{
		ID:                solidutil.TokenID(userID, tokenName),
		UserID:            userID,
		TokenName:         tokenName,
		Token:             solidutil.RandToken(),
		AESKey:            solidutil.RandAESKey(),
		CreatedAtUnixNano: time.Now().UnixNano(),
	}
	core.PanicIf(
		db.Create(&t).Error)
	return &t
}

func (receiver *tokenStore) ListByUserID(userID string) []model.Token {
	var tokens []model.Token
	core.PanicIf(db.Where("user_id = ?", userID).Find(&tokens).Error)
	return tokens
}

func (receiver *tokenStore) Revoke(userID uint, tokenStr string) error {
	return db.Where("user_id = ? AND token = ?", userID, tokenStr).Delete(&model.Token{}).Error
}

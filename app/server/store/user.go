package store

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ismdeep/digest"
	"github.com/ismdeep/log"
	"go.uber.org/zap"

	"github.com/ismdeep/notification/app/server/model"
	"github.com/ismdeep/notification/internal/solidutil"
	"github.com/ismdeep/notification/pkg/core"
)

type userStore struct {
}

// User store
var User *userStore

// GetByID get by id
func (receiver *userStore) GetByID(userID string) *model.User {
	item := &model.User{}
	core.PanicIf(db.Where("id = ?", userID).First(item).Error)
	return item
}

// GetByUsername get by username
func (receiver *userStore) GetByUsername(username string) *model.User {
	item := &model.User{}
	core.PanicIf(db.Where("username = ?", username).First(item).Error)

	return item
}

// ExistsByID check exists by userID
func (receiver *userStore) ExistsByID(userID string) bool {
	var cnt int64
	core.PanicIf(db.Model(&model.User{}).Where("id = ?", userID).Count(&cnt).Error)
	return cnt > 0
}

// ExistsByUsername check exists by username
func (receiver *userStore) ExistsByUsername(username string) bool {
	var cnt int64
	core.PanicIf(db.Model(&model.User{}).Where("username = ?", username).Count(&cnt).Error)
	return cnt > 0
}

func (receiver *userStore) Create(username string, password string) string {
	u := model.User{
		ID:                solidutil.UserID(username),
		Username:          username,
		Nickname:          username,
		Avatar:            "",
		Enabled:           true,
		Digest:            digest.Generate(password),
		CreatedAtUnixNano: time.Now().UnixNano(),
	}

	// 写入用户
	core.PanicIf(db.Create(&u).Error)

	return u.ID
}

func (receiver *userStore) InitUser() {
	var cnt int64
	core.PanicIf(db.Model(&model.User{}).Count(&cnt).Error)

	if cnt > 0 {
		return
	}

	username := "admin"
	password := uuid.NewString()
	adminUserID := receiver.Create(username, password)
	log.WithContext(context.TODO()).Info("success to init user", zap.String("username", username), zap.String("password", password))

	// 自动为admin账号初始化一条 Token
	Token.InitToken(adminUserID)
}

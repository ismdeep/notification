package store

import (
	"time"

	"gorm.io/gorm"

	"github.com/ismdeep/notification/app/server/model"
	"github.com/ismdeep/notification/internal/solidutil"
	"github.com/ismdeep/notification/pkg/core"
)

type msgStore struct {
}

var Msg *msgStore

func (receiver *msgStore) Write(userID string, customerMsgID string, content string) string {
	msgID := solidutil.MsgID(userID, customerMsgID, content)
	core.PanicIf(
		db.Transaction(func(tx *gorm.DB) error {
			var cnt int64
			core.PanicIf(
				db.Model(&model.Msg{}).
					Where("id = ?", msgID).
					Count(&cnt).Error)

			if cnt > 0 {
				return nil
			}

			core.PanicIf(
				db.Create(&model.Msg{
					ID:                msgID,
					CustomerMsgID:     customerMsgID,
					UserID:            userID,
					Content:           content,
					Status:            model.MsgStatusPending,
					Err:               "",
					CreatedAtUnixNano: time.Now().UnixNano(),
				}).Error)

			return nil
		}))
	return msgID
}

func (receiver *msgStore) PendingMsgList(size int) []model.Msg {
	var lst []model.Msg
	core.PanicIf(
		db.Where("status = ?", model.MsgStatusPending).
			Order("created_at_unix_nano ASC").
			Limit(size).
			Find(&lst).Error)
	return lst
}

func (receiver *msgStore) SetSent(id string) {
	core.PanicIf(
		db.Model(&model.Msg{}).
			Where("id = ?", id).
			Update("status", model.MsgStatusSent).
			Error)
}

func (receiver *msgStore) SetFailed(id string, errMsg string) {
	core.PanicIf(
		db.Model(&model.Msg{}).
			Where("id = ?", id).
			Updates(map[string]interface{}{
				"status": model.MsgStatusFailed,
				"err":    errMsg,
			}).Error)
}

package mail

import (
	"github.com/ismdeep/notification/api/model"
	"github.com/ismdeep/notification/api/request"
	userStore "github.com/ismdeep/notification/api/store/user"
	"github.com/ismdeep/notification/common"
	"strings"
)

// PushMail push mail
func PushMail(userID uint, req *request.PushMailRequest) error {
	if req == nil {
		return common.ErrBadRequest
	}

	if err := req.Check(); err != nil {
		return common.ErrBadRequest
	}

	// 判断用户是否存在
	if !userStore.ExistsByID(userID) {
		return common.ErrUserNotExists
	}

	mail := &model.Mail{
		UserID:     userID,
		Status:     model.StatusPending,
		Type:       req.Type,
		Subject:    req.Subject,
		Content:    req.Content,
		ToMailList: strings.Join(req.ToMailList, ";"),
	}
	if err := model.DB.Save(mail).Error; err != nil {
		return common.ErrSystemError
	}

	return nil
}

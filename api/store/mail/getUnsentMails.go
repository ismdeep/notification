package mail

import (
	"errors"
	"github.com/ismdeep/notification/api/model"
	"github.com/ismdeep/notification/common"
	"github.com/jinzhu/gorm"
)

// GetUnsentMails 获取未发送邮件列表
func GetUnsentMails(size uint) ([]*model.Mail, error) {
	if size <= 0 || size > 100 {
		return nil, common.ErrBadRequest
	}

	mails := make([]*model.Mail, 0)

	if err := model.DB.Where("status = ?", model.StatusPending).Limit(size).Find(&mails).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return make([]*model.Mail, 0), nil
		}

		return nil, err
	}

	return mails, nil
}

// MarkedAsSent 标记为已发送
func MarkedAsSent(id uint) error {
	mailItem := &model.Mail{}
	if err := model.DB.Where("id = ?", id).First(mailItem).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("邮件不存在")
		}

		return err
	}

	mailItem.Status = model.StatusSent

	if err := model.DB.Save(mailItem).Error; err != nil {
		return err
	}

	return nil
}

package mail

import (
	"github.com/ismdeep/notification/api/model"
	"github.com/ismdeep/notification/api/response"
	"github.com/jinzhu/gorm"
	"strings"
)

func GetMails(userID uint) ([]*response.MailInfo, error) {
	mails := make([]*model.Mail, 0)
	if err := model.DB.Where("user_id = ?", userID).Find(&mails).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return make([]*response.MailInfo, 0), nil
		}

		return nil, err
	}

	results := make([]*response.MailInfo, 0)
	for _, mailItem := range mails {
		results = append(results, &response.MailInfo{
			ID:         mailItem.ID,
			Status:     mailItem.Status,
			StatusText: "unknown",
			SenderName: mailItem.SenderName,
			Type:       mailItem.Type,
			Content:    mailItem.Content,
			ToMails:    strings.Split(mailItem.ToMailList, ";"),
		})
	}

	return results, nil
}

package mail

import (
	"github.com/ismdeep/notification/api/response"
	mailService "github.com/ismdeep/notification/service/mail"
)

// GetMailTypes get mail types
func GetMailTypes() []*response.MailType {
	results := make([]*response.MailType, 0)
	results = append(results, &response.MailType{Type: mailService.ContentTypeHTML})
	results = append(results, &response.MailType{Type: mailService.ContentTypeText})
	return results
}

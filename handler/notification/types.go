package notification

import (
	"github.com/ismdeep/notification/api/response"
)

// GetTypes get notification types
func GetTypes() []*response.NotificationType {

	return []*response.NotificationType{
		{
			Name: "Email",
			Key:  "email",
		},
		{
			Name: "WeCom Bot",
			Key:  "wecom",
		},
		{
			Name: "Telegram",
			Key:  "telegram",
		},
	}
}

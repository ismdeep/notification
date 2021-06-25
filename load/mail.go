package load

import "github.com/ismdeep/notification/service/mail"

// MailService 加载邮件服务
func MailService() {
	mail.Init()
}

package mail

import (
	"github.com/ismdeep/log"
	"github.com/ismdeep/notification/api/store/mail"
	"github.com/ismdeep/notification/config"
	"gopkg.in/gomail.v2"
	"strings"
	"time"
)

const (
	// ContentTypeHTML ContentType: HTML
	ContentTypeHTML string = "text/html"
	// ContentTypeText ContentType: Text
	ContentTypeText string = "text/plain"
)

var mailSenderService struct {
	Host     string
	Port     int
	Username string
	Password string
}

func daemon() {
	for {
		mails, err := mail.GetUnsentMails(10)
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		if len(mails) <= 0 {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		for _, mailItem := range mails {
			m := gomail.NewMessage()
			m.SetHeader("From", m.FormatAddress(mailSenderService.Username, mailItem.SenderName))
			m.SetHeader("To", strings.Split(mailItem.ToMailList, ";")...)
			m.SetHeader("Subject", mailItem.Subject)
			m.SetBody(mailItem.Type, mailItem.Content)
			d := gomail.NewDialer(mailSenderService.Host, mailSenderService.Port, mailSenderService.Username, mailSenderService.Password)
			if err1 := d.DialAndSend(m); err1 == nil {
				// 发送成功，标记为已发送
				if err := mail.MarkedAsSent(mailItem.ID); err != nil {
					log.Warn("service.mail.daemon", "err", err)
				}

				log.Info("service.mail.daemon", "msg", "email send successfully", "to", mailItem.ToMailList, "subject", mailItem.Subject)
			}

		}
	}
}

// Init 初始化
func Init() {
	mailSenderService.Host = config.Mail.Host
	mailSenderService.Port = config.Mail.Port
	mailSenderService.Username = config.Mail.Username
	mailSenderService.Password = config.Mail.Password

	go func() {
		daemon()
	}()
}

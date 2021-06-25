package mail

import (
	"github.com/ismdeep/log"
	"github.com/ismdeep/notification/config"
	"gopkg.in/gomail.v2"
	"time"
)

const (
	// ContentTypeHTML ContentType: HTML
	ContentTypeHTML string = "text/html"
	// ContentTypeText ContentType: Text
	ContentTypeText string = "text/plain"
)

// Pack pack
type Pack struct {
	SenderName string
	Subject    string
	Type       string
	Content    string
	ToMail     string
}

var mailSenderService struct {
	Host     string
	Port     int
	Username string
	Password string
}

func daemon() {
	for {
		pack := <-packChan

		m := gomail.NewMessage()
		m.SetHeader("From", m.FormatAddress(mailSenderService.Username, pack.SenderName))
		m.SetHeader("To", pack.ToMail)
		m.SetHeader("Subject", pack.Subject)
		m.SetBody(pack.Type, pack.Content)
		d := gomail.NewDialer(mailSenderService.Host, mailSenderService.Port, mailSenderService.Username, mailSenderService.Password)
		if err := d.DialAndSend(m); err != nil {
			log.Error("service.mail.daemon", "err", err, "to", pack.ToMail, "subject", pack.Subject)
			packChan <- pack
			time.Sleep(100 * time.Millisecond)
			continue
		}
		log.Info("service.mail.daemon", "msg", "email send successfully", "to", pack.ToMail, "subject", pack.Subject)
	}
}

var packChan = make(chan *Pack, 1000)

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

// Push 推送邮件
func Push(mailPack *Pack) {
	packChan <- mailPack
}

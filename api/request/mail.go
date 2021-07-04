package request

import (
	"errors"
	"github.com/ismdeep/notification/service/mail"
)

// PushMailRequest push mail request
type PushMailRequest struct {
	SenderName string   `json:"sender_name"`
	Subject    string   `json:"subject"`
	Type       string   `json:"type"`
	Content    string   `json:"content"`
	ToMailList []string `json:"to_mail_list"`
}

// Check check
func (p *PushMailRequest) Check() error {
	if p.SenderName == "" {
		return errors.New("sender_name can NOT be empty")
	}

	if p.Subject == "" {
		return errors.New("subject can NOT be empty")
	}

	if p.Type == "" {
		return errors.New("type can NOT be empty")
	}

	if p.Type != mail.ContentTypeHTML && p.Type != mail.ContentTypeText {
		return errors.New("type error")
	}

	if p.Content == "" {
		return errors.New("content can NOT be empty")
	}

	if len(p.ToMailList) <= 0 {
		return errors.New("to_mail can NOT be empty")
	}

	return nil
}

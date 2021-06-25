package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/log"
	"github.com/ismdeep/notification/config"
	"github.com/ismdeep/notification/service/mail"
)

// PushMailRequest push mail request
type PushMailRequest struct {
	SenderName string `json:"sender_name"`
	Subject    string `json:"subject"`
	Type       string `json:"type"`
	Content    string `json:"content"`
	ToMail     string `json:"to_mail"`
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

	if p.ToMail == "" {
		return errors.New("to_mail can NOT be empty")
	}

	return nil
}

// PushMail push mail
func PushMail(c *gin.Context) {
	authKey := c.GetHeader("Authorization")
	if authKey != fmt.Sprintf("Bearer %v", config.Global.Secret) {
		Error(c, WithMsg("access denied"))
		return
	}

	req := &PushMailRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		Error(c, WithMsg("bind request failed"))
		return
	}

	if err := req.Check(); err != nil {
		Error(c, WithMsg(err.Error()))
		return
	}

	log.Info("PushMail", "req", req)

	mail.Push(&mail.Pack{
		SenderName: req.SenderName,
		Subject:    req.Subject,
		Type:       req.Type,
		Content:    req.Content,
		ToMail:     req.ToMail,
	})

	Ok(c, WithMsg("mail sent"))
	return
}

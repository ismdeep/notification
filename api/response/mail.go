package response

// MailType mail type
type MailType struct {
	Type string `json:"type"`
}

// MailInfo 邮件信息
type MailInfo struct {
	ID         uint     `json:"id"`          // 流水号
	Status     uint     `json:"status"`      // 状态号
	StatusText string   `json:"status_text"` // 状态文本
	SenderName string   `json:"sender_name"` // 抬头
	Type       string   `json:"type"`        // 邮件内容类型
	Content    string   `json:"content"`     // 邮件正文
	ToMails    []string `json:"to_mails"`    // 发送列表
}

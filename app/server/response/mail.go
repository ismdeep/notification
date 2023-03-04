package response

import "time"

// MailType mail type
type MailType struct {
	Type string `json:"type"`
}

// MailInfo 邮件信息
type MailInfo struct {
	ID         uint      `json:"id"`          // 流水号
	Status     uint      `json:"status"`      // 状态号
	Subject    string    `json:"subject"`     // 标题
	SenderName string    `json:"sender_name"` // 抬头
	Content    string    `json:"content"`     // 邮件正文
	Type       string    `json:"type"`        // 邮件内容类型
	ToMails    []string  `json:"to_mails"`    // 发送列表
	FailedMsg  string    `json:"failed_msg"`  // 失败信息
	CreatedAt  time.Time `json:"created_at"`  // 邮件创建时间
}

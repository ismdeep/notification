package request

// PushMailRequest push mail request
type PushMailRequest struct {
	Subject string   `json:"subject"`
	Type    string   `json:"type"`
	Content string   `json:"content"`
	To      []string `json:"to"`
}

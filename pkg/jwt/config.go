package jwt

// Config config struct
type Config struct {
	Key    string `json:"key"`    // 密钥
	Expire string `json:"expire"` // 超时时长，1s,10m
}

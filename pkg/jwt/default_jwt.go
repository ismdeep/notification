package jwt

var defaultJWTClient *JWT

// Init 初始化
func Init(config *Config) {
	defaultJWTClient = New(config)
}

// GenerateToken 生成token
func GenerateToken(content string) (string, error) {
	return defaultJWTClient.GenerateToken(content)
}

// VerifyToken 格式化token
func VerifyToken(tokens string) (string, error) {
	return defaultJWTClient.VerifyToken(tokens)
}

package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWT Struct Info
type JWT struct {
	signKey    []byte        // 密钥
	expireTime time.Duration // 过期时间
}

// claims jwt claims
type claimsStruct struct {
	Content string `json:"content"`
	jwt.MapClaims
}

// New create instance
func New(config *Config) *JWT {
	c := &Config{
		Key:    uuid.NewString(),
		Expire: "72h",
	}
	if config != nil {
		c = config
	}

	if _, err := time.ParseDuration(c.Expire); err != nil {
		c.Expire = "72h"
	}

	instance := &JWT{}
	instance.signKey = []byte(c.Key)
	instance.expireTime, _ = time.ParseDuration(c.Expire)

	return instance
}

// GenerateToken generate token
func (receiver *JWT) GenerateToken(content string) (token string, err error) {
	claim := &claimsStruct{
		Content: content,
		MapClaims: jwt.MapClaims{
			"ExpiresAt": time.Now().Add(receiver.expireTime).Unix(),
		},
	}
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	token, err = tokens.SignedString(receiver.signKey)
	return token, err
}

// VerifyToken verify token
func (receiver *JWT) VerifyToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, receiver.secret())
	if err != nil {
		return "", err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to map claim")
		return "", err
	}
	if !token.Valid {
		return "", errors.New("token is invalid")
	}

	return claim["content"].(string), nil
}

func (receiver *JWT) secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return receiver.signKey, nil
	}
}

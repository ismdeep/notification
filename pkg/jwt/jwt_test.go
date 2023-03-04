package jwt

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJWT_VerifyToken(t *testing.T) {
	j := New(&Config{
		Key:    uuid.NewString(),
		Expire: "72h",
	})
	token, err := j.GenerateToken("hello")
	assert.NoError(t, err)
	_, err1 := jwt.Parse(token, j.secret())
	assert.NoError(t, err1)

	content, err2 := j.VerifyToken(token)
	assert.NoError(t, err2)

	t.Logf("content = %v", content)
}

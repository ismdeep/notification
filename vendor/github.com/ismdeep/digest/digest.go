package digest

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"reflect"
)

const saltSize = 64

func sha512WithSalt(password string, salt []byte) []byte {
	stream := make([]byte, 0)
	for _, v := range []byte(password) {
		stream = append(stream, v)
	}
	for _, v := range salt {
		stream = append(stream, v)
	}

	s := sha512.New()
	s.Write(stream)
	r := s.Sum(nil)
	return r
}

// Generate 生成安全摘要 salt := randBytes(64);  base64(salt + sha512(password + salt))
func Generate(password string) string {
	salt := make([]byte, saltSize)
	_, _ = rand.Read(salt)

	r := sha512WithSalt(password, salt)

	resultBytes := make([]byte, 0)
	for _, v := range salt {
		resultBytes = append(resultBytes, v)
	}
	for _, v := range r {
		resultBytes = append(resultBytes, v)
	}

	result := base64.StdEncoding.EncodeToString(resultBytes)

	return result
}

func Verify(digest string, password string) bool {
	decodeBytes, err := base64.StdEncoding.DecodeString(digest)
	if err != nil {
		return false
	}

	if len(decodeBytes) != saltSize+64 {
		return false
	}

	salt := make([]byte, 0)
	for i := 0; i < saltSize; i++ {
		salt = append(salt, decodeBytes[i])
	}

	hashBytes := make([]byte, 0)
	for i := saltSize; i < 64+saltSize; i++ {
		hashBytes = append(hashBytes, decodeBytes[i])
	}

	// sha512(password + salt) == hashBytes
	r := sha512WithSalt(password, salt)

	return reflect.DeepEqual(r, hashBytes)
}

package handler

import (
	"testing"

	"github.com/google/uuid"

	"github.com/ismdeep/notification/app/server/request"
	"github.com/ismdeep/notification/app/server/store"
)

func Test_tokenHandler_Create(t *testing.T) {
	t.Logf("got = %v",
		Token.Create(store.User.GetByUsername("admin").ID, request.NewToken{
			Name: "test-token-" + uuid.NewString(),
		}).Token)
}

package handler

import (
	"errors"

	"github.com/ismdeep/notification/app/server/request"
	"github.com/ismdeep/notification/app/server/response"
	"github.com/ismdeep/notification/app/server/store"
	"github.com/ismdeep/notification/internal/solidutil"
	"github.com/ismdeep/notification/pkg/core"
)

type tokenHandler struct {
}

// Token handler
var Token *tokenHandler

// Create new token
func (receiver *tokenHandler) Create(userID string, req request.NewToken) *response.TokenDetail {
	core.PanicIf(
		core.IfErr(req.Name == "", errors.New("bad request")))

	store.User.GetByID(userID)

	core.PanicIf(
		core.IfErr(
			store.Token.ExistsByID(solidutil.TokenID(userID, req.Name)), errors.New("already exists")))

	token := store.Token.Create(userID, req.Name)

	return &response.TokenDetail{
		Name:  req.Name,
		Token: token.Token,
	}
}

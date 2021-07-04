package token

import (
	"errors"
	"github.com/ismdeep/notification/api/model"
	"github.com/ismdeep/notification/api/request"
	"github.com/ismdeep/notification/api/response"
	"github.com/ismdeep/notification/common"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"log"
)

// NewToken new token
func NewToken(userID uint, req *request.NewToken) (*response.TokenDetail, error) {
	if req == nil {
		return nil, common.ErrBadRequest
	}

	if req.Name == "" {
		return nil, common.ErrBadRequest
	}

	user := &model.User{}
	if err := model.DB.Where("id = ?", userID).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrUserNotExists
		}

		return nil, common.ErrSystemError
	}

	if !user.Enabled {
		return nil, common.ErrUserHasBeenBaned
	}

	token := &model.Token{
		UserID:    userID,
		TokenName: req.Name,
		Token:     uuid.NewV4().String(),
	}

	if err := model.DB.Save(token).Error; err != nil {
		log.Println(err)
		return nil, errors.New("添加失败")
	}

	return &response.TokenDetail{
		Name:  req.Name,
		Token: token.Token,
	}, nil
}

package token

import (
	"github.com/ismdeep/notification/api/model"
	"github.com/ismdeep/notification/api/response"
	"github.com/ismdeep/notification/common"
	"github.com/jinzhu/gorm"
)

// List list all tokens
func List(userID uint) ([]*response.TokenDetail, error) {
	user := &model.User{}
	if err := model.DB.Where("id = ?", userID).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrUserNotExists
		}

		return nil, common.ErrSystemError
	}

	tokens := make([]*model.Token, 0)

	if err := model.DB.Where("user_id = ?", userID).Find(&tokens).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return make([]*response.TokenDetail, 0), nil
		}

		return nil, err
	}

	results := make([]*response.TokenDetail, 0)
	for _, token := range tokens {
		results = append(results, &response.TokenDetail{
			Name:  token.TokenName,
			Token: token.Token,
		})
	}

	return results, nil
}

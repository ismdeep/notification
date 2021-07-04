package token

import (
	"errors"
	"github.com/ismdeep/notification/api/model"
	"github.com/ismdeep/notification/common"
	"github.com/jinzhu/gorm"
)

// Revoke revoke token
func Revoke(userID uint, tokenStr string) error {
	user := &model.User{}
	if err := model.DB.Where("id = ?", userID).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrUserNotExists
		}

		return common.ErrSystemError
	}

	if tokenStr == "" {
		return common.ErrBadRequest
	}

	where := model.DB.Where("user_id = ? and token = ?", userID, tokenStr)

	token := &model.Token{}

	if err := where.First(token).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("不存在的Token")
		}

		return common.ErrSystemError
	}

	if err := where.Delete(token).Error; err != nil {
		return common.ErrSystemError
	}

	return nil
}

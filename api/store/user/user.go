package user

import (
	"github.com/ismdeep/notification/api/model"
	"github.com/jinzhu/gorm"
)

// GetByID get by id
func GetByID(userID uint) (*model.User, error) {
	item := &model.User{}
	if err := model.DB.Where("id = ?", userID).First(item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// GetByUsername get by username
func GetByUsername(username string) (*model.User, error) {
	item := &model.User{}
	if err := model.DB.Where("username = ?", username).First(item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

// ExistsByID check exists by userID
func ExistsByID(userID uint) bool {
	item, _ := GetByID(userID)
	return item != nil
}

// ExistsByUsername check exists by username
func ExistsByUsername(username string) bool {
	item, _ := GetByUsername(username)
	return item != nil
}

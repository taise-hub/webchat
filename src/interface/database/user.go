package database

import (
	"gorm.io/gorm"
	"github.com/taise-hub/webchat/src/domain/model"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository {
		db: db,
	}
}

func (rep *userRepository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	result := rep.db.Find(user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (rep *userRepository) GetByID(id uint) (*model.User, error) {
	user := &model.User{}
	result := rep.db.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (rep *userRepository) Create(user *model.User) error {
	result := rep.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
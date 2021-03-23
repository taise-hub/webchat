package database

import (
	"gorm.io/gorm"
	"github.com/taise-hub/webchat/src/domain/model"
)

type UserRepository interface {
	GetByEmail(string) (*model.User, error)
	Create(*model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
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

func (rep *userRepository) Create(user *model.User) error {
	result := rep.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
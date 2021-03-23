package database

import (
	"gorm.io/gorm"
	"github.com/taise-hub/webchat/src/domain/model"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository {
		db: db,
	}
}

func (rep *MessageRepository) Save(message *model.Message) error {
	result := rep.db.Create(message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
package database

import (
	"gorm.io/gorm"
	"github.com/taise-hub/webchat/src/domain/model"
)

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository {
		db: db,
	}
}

func (rep *messageRepository) Save(message *model.Message) error {
	result := rep.db.Create(message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (rep *messageRepository) GetAll() (*model.Messages, error) {
	var messages model.Messages
	result := rep.db.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return &messages, nil
}
package controller

import (
	"fmt"
	"github.com/taise-hub/webchat/src/domain/model"
	"github.com/taise-hub/webchat/src/usecase"
	"github.com/taise-hub/webchat/src/interface/database"
	"gorm.io/gorm"
)

type MessageController interface {
	Save(string, uint) bool
	GetAll() (*[]model.Message, error)
}

type messageController struct {
	usecase usecase.MessageUsecase
}

func NewMessageController(db *gorm.DB) *messageController {
	return &messageController {
		usecase: usecase.MessageUsecase {
			Repository: database.NewMessageRepository(db),
		},
	}
}

func (con *messageController) Save(text string, userID uint) bool {
	if err := con.usecase.Save(text, userID); err != nil {
		fmt.Printf("%+v", err)
		return false
	}
	return true
}

func (con *messageController) GetAll() (*[]model.Message, error) {
	messages, err := con.usecase.GetAll()
	if err != nil {
		return nil, err
	}
	return messages, nil
}
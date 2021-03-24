package usecase

import (
	"github.com/taise-hub/webchat/src/domain/model"
	"github.com/taise-hub/webchat/src/domain/repository"
)

type MessageUsecase struct {
	Repository repository.MessageRepository
}

func (uc *MessageUsecase) Save(text string, userID uint) error {
	message := &model.Message{
		Text: text,
		UserID: userID,
	}
	if err := uc.Repository.Save(message); err != nil {
		return err
	}
	return nil
}

func (uc *MessageUsecase) GetAll() (*model.Messages, error) {
	messages, err := uc.Repository.GetAll()
	if err != nil {
		return nil, err
	}
	return messages, nil
}
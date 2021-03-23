package usecase

import (
	"github.com/taise-hub/webchat/src/domain/model"
	"github.com/taise-hub/webchat/src/domain/repository"
)

type MessageUsecase struct {
	Repository repository.MessageRepository
}

func (uc *MessageUsecase) Save(text string, user *model.User) error {
	message := &model.Message{
		Id: 1,
		Text: text,
		User: user,
	}
	if err := uc.Repository.Save(message); err != nil {
		return err
	}
	return nil
}
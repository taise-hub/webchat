package repository

import (
	"github.com/taise-hub/webchat/src/domain/model"
)

type MessageRepository interface{
	Save(*model.Message) error
}
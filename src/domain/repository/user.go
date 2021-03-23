package repository

import (
	"github.com/taise-hub/webchat/src/domain/model"
)


type UserRepository interface {
	GetByEmail(string) (*model.User, error)
	Create(*model.User) error
}


package usecase


import (
	"github.com/taise-hub/webchat/src/domain/model"
	"github.com/taise-hub/webchat/src/domain/repository"
)


type userUsecase struct {
	rep repository.UserRepository
}

func (us *userUsecase) GetByEmail(email string) (*model.User, error) {
	user, err := us.rep.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userUsecase) Create(user *model.User) error {
	 if err := us.rep.Create(user); err != nil {
		 return err
	 }
	 return nil
}
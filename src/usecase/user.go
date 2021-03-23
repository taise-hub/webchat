package usecase


import (
	"github.com/taise-hub/webchat/src/domain/model"
	"github.com/taise-hub/webchat/src/domain/repository"
)

type UserUsecase struct {
	Repository repository.UserRepository
}


func (us *UserUsecase) GetByEmail(email string) (*model.User, error) {
	user, err := us.Repository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserUsecase) Create(user *model.User) error {
	 if err := us.Repository.Create(user); err != nil {
		 return err
	 }
	 return nil
}
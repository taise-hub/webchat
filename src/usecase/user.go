package usecase


import (
	"fmt"
	"crypto/sha256"
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

func (us *UserUsecase) Create(email string, name string, password string) error {
	sha256pass := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	user := &model.User{
		Email: email,
		Name: name,
		Password: sha256pass,
	}
	 if err := us.Repository.Create(user); err != nil {
		 return err
	 }
	 return nil
}

func (us *UserUsecase) Login(email string, password string) (bool, error) {
	user, err := us.GetByEmail(email)
	if err != nil {
		return false, err
	}
	dbPass := user.Password
	formPass := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	if dbPass != formPass {
		return false, nil
	}
	return true, nil
}
package controller


import (
	"github.com/taise-hub/webchat/src/domain/model"
	"github.com/taise-hub/webchat/src/usecase"
	"github.com/taise-hub/webchat/src/interface/database"
	"gorm.io/gorm"
)

type UserController struct {
	Usecase usecase.UserUsecase
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		Usecase: usecase.UserUsecase {
			Repository: database.NewUserRepository(db),
		},
	}
}

func (con *UserController) GetByEmail(email string) (*model.User, error) {
	user, err :=con.Usecase.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (con *UserController) Create(email string, name string, password string) error {
	if err := con.Usecase.Create(email, name, password); err != nil {
		return err
	}
	return nil
}
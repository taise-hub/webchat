package controller


import (
	"gorm.io/gorm"
	"github.com/taise-hub/webchat/src/domain/model"
	"github.com/taise-hub/webchat/src/usecase"
	"github.com/taise-hub/webchat/src/interface/database"
)

type userController struct {
	Usecase usecase.UserUsecase
}

func NewUserController(db *gorm.DB) *userController {
	return &userController{
		Usecase: usecase.UserUsecase {
			Repository: database.NewUserRepository(db),
		},
	}
}

func (con *userController) GetByEmail(email string) (*model.User, error) {
	user, err :=con.Usecase.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (con *userController) Create(email string, name string, password string) {
	user := &model.User {
		Email: email,
		Name: name,
		Password: password,
	}
	err := con.Usecase.Create(user)
	if err != nil {
		panic(err) //TODO:正しいerrorハンドリングを考える
	}
}
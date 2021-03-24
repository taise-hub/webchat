package controller

import (
	"fmt"
	"github.com/taise-hub/webchat/src/usecase"
)

type messageController struct {
	usecase usecase.MessageUsecase
}

func(con *messageController) Save(text string, userID uint) bool {
	if err := con.usecase.Save(text, userID); err != nil {
		fmt.Printf("%+v", err)
		return false
	}
	return true
}
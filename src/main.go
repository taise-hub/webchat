package main

import (
	"github.com/taise-hub/webchat/src/domain/model"
	"github.com/taise-hub/webchat/src/infrastructure"
	"github.com/taise-hub/webchat/src/router"
)

func main() {
	db, err := infrastructure.NewDB(infrastructure.NewConfig())
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Message{})
	if err != nil {
		panic(err)
	}
	router.Init(db)
}
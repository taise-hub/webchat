package main

import (
	"github.com/taise-hub/webchat/src/domain/model"
	"github.com/taise-hub/webchat/src/infrastructure"
	"github.com/taise-hub/webchat/src/router"
)

func main() {
	db, err := infrastructure.NewDB(infrastructure.NewConfig())
	db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	router.Init(db)
}
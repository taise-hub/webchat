package main

import (
	"github.com/taise-hub/webchat/src/infrastructure"
	"github.com/taise-hub/webchat/src/interface/controller"
)

func main() {
	db, err := infrastructure.NewDB(infrastructure.NewConfig())
	if err != nil {
		panic(err)
	}
	controller.Init(db)
}
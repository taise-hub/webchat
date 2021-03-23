package controller

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func Init(db *gorm.DB) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	userController := NewUserController(db)

	router.GET("/signup", userController.GetSignUp)
	router.POST("/signup", userController.PostSignUp)
	router.GET("/login", userController.GetLogin)

	router.Run(":8080")
}
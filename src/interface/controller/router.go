package controller

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func Init(db *gorm.DB) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	userController := NewUserController(db)
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("SSID", store))

	router.GET("/signup", userController.GetSignUp)
	router.POST("/signup", userController.PostSignUp)
	router.GET("/login", userController.GetLogin)
	router.POST("/login", userController.PostLogin)

	router.Run(":8080")
}
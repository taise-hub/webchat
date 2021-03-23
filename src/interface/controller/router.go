package controller

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func Init(db *gorm.DB) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	// userController := NewUserController(db)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"hello": "world",
		})
	})

	router.Run(":8080")
}
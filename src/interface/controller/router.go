package controller

import (
	"github.com/taise-hub/webchat/src/usecase"
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
	hub := usecase.NewHub()
	go hub.Run()

	router.GET("/signup", userController.GetSignUp)
	router.POST("/signup", userController.PostSignUp)
	router.GET("/login", userController.GetLogin)
	router.POST("/login", userController.PostLogin)

	//service
	service := router.Group("/home")
	service.Use(LoginCheck)
	{
		service.GET("/", userController.GetHome)
		service.POST("/logout", userController.Logout)
		//web socket chat
		service.GET("/chat", userController.GetChat)
		service.GET("/chat/ws", func(c *gin.Context) {
			userController.WsChat(c, hub)
		})
	}

	router.Run(":8080")
}

//TODO:middlewareの正しい作りかたを調べる。
func LoginCheck(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("email")
	if user == nil {
		c.Redirect(302, "/login")
	}
}

package router

import (
	"github.com/taise-hub/webchat/src/chat"
	"github.com/taise-hub/webchat/src/interface/controller"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func Init(db *gorm.DB) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	userController := controller.NewUserController(db)
	messageController := controller.NewMessageController(db)
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("SSID", store))
	hub := chat.NewHub()
	go hub.Run()

	router.GET("/signup", GetSignUp)
	router.POST("/signup", func(c *gin.Context) {
		PostSignUp(c, userController)
	})
	router.GET("/login", GetLogin)
	router.POST("/login", func(c *gin.Context) {
		PostLogin(c, userController)
	})

	//service
	service := router.Group("/home")
	service.Use(LoginCheck)
	{
		service.GET("/", func(c *gin.Context) {
			GetHome(c, userController)
		})
		service.POST("/logout", Logout)
		//web socket chat
		service.GET("/chat", func(c *gin.Context) {
			GetChat(c, userController, messageController)
		})
		service.GET("/chat/ws", func(c *gin.Context) {
			WsChat(c, userController, messageController, hub)
		})
		//[TEST]get messages
		service.GET("/messages", func(c *gin.Context) {
			msgs, err := messageController.GetAll()
			if err != nil {
				c.JSON(500, err)
				return
			}
			c.JSON(200, msgs)
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

package router

import (
	"fmt"
	"github.com/taise-hub/webchat/src/chat"
	"github.com/taise-hub/webchat/src/interface/controller"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

func GetSignUp(c *gin.Context) {
	c.HTML(200, "signup.html", nil)
} 

func PostSignUp(c *gin.Context, uCon *controller.UserController) {
	email := c.PostForm("email")
	name := c.PostForm("name")
	password := c.PostForm("password")
	err := uCon.Create(email, name, password)
	if err != nil {
		c.HTML(400, "signup.html", gin.H{"err": "すでに登録されたメールアドレスです"})
		return
	}
	c.Redirect(302, "/login")
}

func GetLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func PostLogin(c *gin.Context, uCon *controller.UserController) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	ok, err := uCon.Usecase.Login(email, password)
	if err != nil {
		c.JSON(500, "Internal Server Error")
		return
	}
	if !ok {
		c.HTML(400, "login.html", nil)
		return
	}
	session := sessions.Default(c)
	session.Set("email", email)
	session.Save()
	c.Redirect(302, "/home")
}

func GetHome(c *gin.Context, uCon *controller.UserController) {
	session := sessions.Default(c)
	email := session.Get("email")
	user, _ := uCon.GetByEmail(email.(string))
	c.HTML(200, "home.html", gin.H{"user": user})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(302, "/login")
}

func GetChat(c *gin.Context, uCon *controller.UserController, mCon controller.MessageController) {
	msgs, err := mCon.GetAll()
	if err != nil {
		c.JSON(500, err)
		return
	}
	var messages []string
	for _, msg := range *msgs {
		user, err := uCon.GetByID(msg.UserID)
		if err != nil {
			c.JSON(500, err)
			return
		}
		message := fmt.Sprintf("【%s】: %s", user.Name, msg.Text)
		messages = append(messages, message)
	}
	c.HTML(200, "chat.html", gin.H{"messages": messages})
}

func WsChat(c *gin.Context, uCon *controller.UserController, mCon controller.MessageController, hub *chat.Hub) {
	conn, err := chat.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &chat.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client
	//clientがconn.ReadMessage()したら、hubに通知して各clientに流し込む.
	session := sessions.Default(c)
	user, _ := uCon.GetByEmail(session.Get("email").(string))
	go client.Listen(user, mCon)
	go client.Write()
}
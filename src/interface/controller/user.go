package controller


import (
	"fmt"
	"github.com/taise-hub/webchat/src/domain/model"
	"github.com/taise-hub/webchat/src/usecase"
	"github.com/taise-hub/webchat/src/interface/database"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

type userController struct {
	Usecase usecase.UserUsecase
}

func NewUserController(db *gorm.DB) *userController {
	return &userController{
		Usecase: usecase.UserUsecase {
			Repository: database.NewUserRepository(db),
		},
	}
}

func (con *userController) getByEmail(email string) (*model.User, error) {
	user, err :=con.Usecase.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (con *userController) GetSignUp(c *gin.Context) {
	c.HTML(200, "signup.html", nil)
} 

func (con *userController) PostSignUp(c *gin.Context) {
	email := c.PostForm("email")
	name := c.PostForm("name")
	password := c.PostForm("password")
	err := con.Usecase.Create(email, name, password)
	if err != nil {
		c.HTML(400, "signup.html", gin.H{"err": "すでに登録されたメールアドレスです"})
		return
	}
	c.Redirect(302, "/login")
}

func (con *userController) GetLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func (con *userController) PostLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	ok, err := con.Usecase.Login(email, password)
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

func (con *userController) GetHome(c *gin.Context) {
	session := sessions.Default(c)
	email := session.Get("email")
	user, _ := con.getByEmail(email.(string))
	c.HTML(200, "home.html", gin.H{"user": user})
}

func (con *userController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(302, "/login")
}

func (con *userController) GetChat(c *gin.Context) {
	c.HTML(200, "chat.html", nil)
}

func (con *userController) WsChat(c *gin.Context, hub *usecase.Hub) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &usecase.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client
	//clientがconn.ReadMessage()したら、hubに通知して各clientに流し込む.
	session := sessions.Default(c)
	user, _ := con.getByEmail(session.Get("email").(string))
	go client.Listen(user.Name)
	go client.Write()
	
}
package chat

import(
	"fmt"
	"github.com/taise-hub/webchat/src/domain/model"
	"github.com/taise-hub/webchat/src/interface/controller"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

type Client struct{
	Hub *Hub
	Conn *websocket.Conn
	Send chan []byte
}

//あるクライアントからのメッセージをListenしHubに流す。
func (c *Client) Listen(user *model.User, msgCon controller.MessageController) {
	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		text := fmt.Sprintf("%s", p)
		if ok := msgCon.Save(text, user.ID); !ok {
			return
		}
		msg := fmt.Sprintf("【%s】: %s", user.Name, p)
		p = []byte(msg)
		c.Hub.Broadcast <- p
	}
}

func (c *Client) Write() {
	for {
		select {
		case msg := <- c.Send:
			if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
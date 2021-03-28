package chat

import(
	"fmt"
	"github.com/taise-hub/webchat/src/domain/model"
	"github.com/taise-hub/webchat/src/imakita"
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
		if text == "今北産業" {
			msg := fmt.Sprintf("【%s】: %s", user.Name, p)
			p = []byte(msg)
			c.Hub.Broadcast <- p
			msgs, err := msgCon.GetAll()
			if err != nil {
				fmt.Println(err)
				return
			}
			sentence := ""
			for _, msg := range *msgs {
				sentence += msg.Text + ", "
			}
			imakitaController(c, sentence)
			continue
		}
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

func imakitaController(c *Client, sentence string) {
	imakita, err := imakita.Imakita(sentence)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	for i, s := range imakita {
		s = "【今北産業】: " + s
		p := []byte(s)
		c.Hub.Broadcast <- p
		if i == 2 {
			return
		}
	}
} 
package chat

import(
	"fmt"
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
func (c *Client) Listen(name string) {
	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := fmt.Sprintf("【%s】: %s", name, p)
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
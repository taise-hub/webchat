package chat

type Hub struct {
	Clients map[*Client]bool
	Broadcast chan []byte
	Register chan *Client
}



func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast: 	make(chan []byte),
		Register: 	make(chan *Client),
		}
}

//hubの起動
//あるクライアントが受け取ったメッセージを他のクライアントにブロードキャストする
func (h *Hub) Run() {
	for {
		select {
		case client := <- h.Register:
			h.Clients[client] = true
		case msg := <- h.Broadcast:
			for client := range h.Clients {
				client.Send <- msg
			}
		}
	}
}

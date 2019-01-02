package structs

type Hub struct {
	broadcast chan []byte // Канал сообщений
	register chan *Client // Канал регистрации клиентов
	unregister chan *Client //Канал отписки клиентов
	clients map[*Client]bool // Мапа, где ключи - клиенты, значения - BOOL
}

func newHub() *Hub {
	return &Hub{
		broadcast: make(chan []byte),
		register: make(chan *Client),
		unregister: make(chan *Client),
		clients: make(map[*Client]bool),
	}
}

func (hub *Hub) Run () {
	for {
		select {
		case client:=<-hub.register :
			hub.clients[client] = true
		case client:= <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.Send)
			}
		case message := <- hub.broadcast :
			for client := range hub.clients {
				select {
				case client.Send <-message:
				default:
					close(client.Send)
					delete(hub.clients, client)
				}
			}
		}
	}
}
package structs

import "github.com/gorilla/websocket"


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	Hub *Hub
	Conn *websocket.Conn
	Send chan []byte
}




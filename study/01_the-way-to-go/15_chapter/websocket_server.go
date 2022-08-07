package main

import (
	"github.com/gorilla/websocket"
)

func server(ws *websocket.Conn)  {
	ws.ReadMessage()
}
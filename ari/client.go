package ari

import (
	"github.com/gorilla/websocket"
)

func ConnectARI() (*websocket.Conn, error) {
	url := "ws://localhost:8088/ari/events?api_key=admin:secret&app=askai"
	return websocket.DefaultDialer.Dial(url, nil)
}
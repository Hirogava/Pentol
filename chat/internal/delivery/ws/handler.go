package ws

import (
	"log"
	"net/http"

	"github.com/Hirogava/pentol/internal/domain/message"
)

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan message.Message, 256),
	}

	hub.register <- client

	go client.writePump()

	go client.readPump()
}
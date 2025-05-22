package ws

import (
	"log"

	"github.com/Hirogava/pentol/internal/domain/message"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan message.Message
	register   chan *Client
	unregister chan *Client
	history    []message.Message
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan message.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		history:    []message.Message{},
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			log.Printf("Client connected")

			for _, msg := range h.history {
				client.send <- msg
				if len(h.history) == 0 {
					break
				}
			}
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			log.Printf("Client disconnected")
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
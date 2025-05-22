package ws

import (
	"context"
	"log"

	"github.com/Hirogava/pentol/internal/cache"
	"github.com/Hirogava/pentol/internal/domain/message"
)

type Hub struct {
	clients    map[*Client]bool
	incoming   chan message.Message
	outgoing   chan message.Message
	register   chan *Client
	unregister chan *Client
	pubsub     *redis.PubSub
}

func NewHub(pub *redis.PubSub) *Hub {
	return &Hub{
		incoming:   make(chan message.Message),
		outgoing:   make(chan message.Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		pubsub: 	pub,
	}
}

func (h *Hub) Run(ctx context.Context) {
	go h.pubsub.Subscribe(ctx, h.incoming)

	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			log.Printf("Client connected")
			
			history, err := h.pubsub.History(ctx, 50)
			if err != nil {
				log.Println("⚠️ Error loading history:", err)
			} else {
				for _, msg := range history {
					client.send <- msg
				}
			}
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			log.Printf("Client disconnected")
		case message := <-h.incoming:
			if message.SenderID == "redis" {
				continue
			}
			log.Printf("Received message: %s", message)

			if err := h.pubsub.Publish(ctx, message); err != nil {
				log.Println("❌ Redis publish error:", err)
			}

			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case <-ctx.Done():
			log.Println("🛑 Hub shutting down")
			return
		}
	}
}

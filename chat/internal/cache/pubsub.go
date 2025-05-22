package redis

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Hirogava/pentol/internal/domain/message"
	"github.com/redis/go-redis/v9"
)

const channelName = "chat"

type PubSub struct {
	client *redis.Client
}

func NewPubSub(addr string) *PubSub {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &PubSub{client: rdb}
}

func (ps *PubSub) Publish(ctx context.Context, msg message.Message) error {
	msg.SenderID = "redis"
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if err := ps.client.LPush(ctx, "chat_history", data).Err(); err != nil {
		log.Println("⚠️ Error writing to Redis history:", err)
	}

	return ps.client.Publish(ctx, channelName, data).Err()
}

func (ps *PubSub) Subscribe(ctx context.Context, outgoing chan<- message.Message) {
	sub := ps.client.Subscribe(ctx, channelName)
	ch := sub.Channel()

	for {
		select {
		case <-ctx.Done():
			log.Println("🔌 Redis subscriber stopped")
			return
		case msg := <-ch:
			var m message.Message
			if err := json.Unmarshal([]byte(msg.Payload), &m); err != nil {
				log.Println("⚠️ Error unmarshaling from Redis:", err)
				continue
			}
			outgoing <- m
		}
	}
}

func (ps *PubSub) History(ctx context.Context, limit int64) ([]message.Message, error) {
	data, err := ps.client.LRange(ctx, "chat_history", 0, limit-1).Result()
	if err != nil {
		return nil, err
	}

	history := make([]message.Message, 0, len(data))
	for _, item := range data {
		var msg message.Message
		if err := json.Unmarshal([]byte(item), &msg); err == nil {
			msg.SenderID = ""
			history = append(history, msg)
		}
	}

	return history, nil
}
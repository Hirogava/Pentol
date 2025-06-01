package channel

import "time"

type Channel struct {
	Id      int `json:"id"`
	OwnerId int `json:"owner_id"`
}

type ChannelDesc struct {
	ChannelData Channel `json:"channel"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
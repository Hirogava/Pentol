package message

import "time"

type Message struct {
    User string     `json:"user"`
    Text string     `json:"text"`
    TS   time.Time  `json:"ts"`
    SenderID string `json:"sender_id"`
}
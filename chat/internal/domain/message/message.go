package message

import "time"

type Message struct {
    User string     `json:"user"`
    Text string     `json:"text"`
    TS   time.Time  `json:"ts"`
    SenderID string `json:"sender_id"`
    UserUID string  `json:"user_uid"`
}

type MessageNew struct {
    Id int    `json:"id"`
    ChatID int `json:"chat_id"`
    SenderID int `json:"sender_id"`
    TS   time.Time  `json:"ts"`
    Text string `json:"text"`
}
package main

import "time"

type Message struct {
	UserID    string `json:"userId"`
	ChatID    string `json:"chatId"`
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
}

func NewMessage(userID string, chatID string, text string) *Message {
	return &Message{userID, chatID, text, time.Now().Unix()}
}

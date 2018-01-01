package main

type Message struct {
	UserID string `json:"userId"`
	Text   string `json:"text"`
}

func NewMessage(userID string, chatID string, text string) *Message {
	return &Message{userID, text}
}

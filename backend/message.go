package main

type Message struct {
	Author string
	Text   string
}

func NewMessage(author string, text string) *Message {
	return &Message{author, text}
}

package main

import (
	"time"
)

type Chat struct {
	Name      string
	Broadcast chan Packet
	Users     []UserConn
}

type Message struct {
	User      int
	Timestamp time.Time
	Chat      *Chat
	Text      string
}

func NewChat(name string) *Chat {
	return &Chat{
		name,
		make(chan Packet),
		[]UserConn{},
	}
}

func NewMessage(u User, t string) *Message {
	return &Message{
		u.ID,
		time.Now(),
		nil,
		t,
	}
}

func (c *Chat) Write(r *Router) {
	for p := range c.Broadcast {
		for _, u := range c.Users {
			u.Conn.WriteJSON(p)
		}
	}
}

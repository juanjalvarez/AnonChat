package main

import "github.com/gorilla/websocket"

type Session struct {
	User *User
	Conn *websocket.Conn
	Chat *Chat
}

func NewSession(conn *websocket.Conn) (*Session, error) {
	u, err := NewUser()
	if err != nil {
		return nil, err
	}
	ss := Session{
		u,
		conn,
		nil,
	}
	return &ss, nil
}

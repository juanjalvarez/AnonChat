package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Session struct {
	User *User
	Conn *websocket.Conn
	Send chan *Event
}

func NewSession(conn *websocket.Conn) *Session {
	ss := Session{
		nil,
		conn,
		make(chan *Event),
	}
	return &ss
}

func (ss *Session) Read(s *Server) {
	var e Event
	for {
		if err := ss.Conn.ReadJSON(&e); err != nil {
			fmt.Println(err)
			break
		}
		s.HandleEvent(ss, &e)
	}
	s.EndSession(ss)
}

func (ss *Session) Write(s *Server) {
	var err error
	for e := range ss.Send {
		err = ss.Conn.WriteJSON(e)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func (ss *Session) Authenticate(u *User) {
	ss.User = u
}

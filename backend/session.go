package main

import (
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
	var err error
	for {
		if err != nil {
			break
		}
		if err = ss.Conn.ReadJSON(&e); err != nil {
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
			break
		}
	}
	s.EndSession(ss)
}

func (ss *Session) Authenticate(u *User) {
	ss.User = u
}

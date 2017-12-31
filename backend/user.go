package main

import (
	"fmt"
)

type User struct {
	ID   string
	Name string
}

func NewUser() (*User, error) {
	newId, err := generateKey(8)
	if err != nil {
		return nil, err
	}
	return &User{
		string(newId),
		"anonymous",
	}, nil
}

func (ss *Session) Read(s *Server) {
	var e *Event
	for {
		if err := ss.Conn.ReadJSON(e); err != nil {
			fmt.Println(err)
			break
		}
		s.HandleEvent(ss, e)
	}
	s.EndSession(ss)
}

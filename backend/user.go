package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type UserConn struct {
	User *User
	Conn *websocket.Conn
	Chat *Chat
}

type User struct {
	ID   int
	Name string
}

func NewUserConn(conn *websocket.Conn) *UserConn {
	u := &User{
		1,
		"anonymous",
	}
	uc := UserConn{
		u,
		conn,
		nil,
	}
	return &uc
}

func (uc *UserConn) Read(r *Router) {
	var p *Packet
	for {
		if err := uc.Conn.ReadJSON(p); err != nil {
			fmt.Println(err)
			break
		}
		r.RoutePacket(uc, p)
	}
	r.DisconnectUser(uc.User.ID)
}

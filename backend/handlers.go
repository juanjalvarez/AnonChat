package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type EventHandler func(*Server, *Session, *Event)

type NewChatRequest struct {
	Name string `json:"name"`
}

type SetUserRequest struct {
	UserID int    `json:"userId"`
	Name   string `json:"name"`
}

type NewMessageRequest struct {
	Text   string `json:"text"`
	ChatID string `json:"chatId"`
}

type ChatSubscribeRequest struct {
	ChatID string `json:"chatId"`
}

type AuthenticationRequest struct {
	NewUser bool   `json:"newUser"`
	Token   string `json:"token"`
}

type AuthenticationResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func registerHandlers(s *Server) {
	s.Handle("new_chat", handleNewChat)
	s.Handle("new_message", handleNewMessage)
	s.Handle("set_user", handleSetUser)
	s.Handle("subscribe_chat", handleSubscribeChat)
	s.Handle("authenticate", handleAuthentication)
}

func handleNewChat(s *Server, ss *Session, e *Event) {
	var req NewChatRequest
	if err := mapstructure.Decode(e.Data, &req); err != nil {
		fmt.Println(err)
		return
	}
	nc, err := NewChat(req.Name, ss.User)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.NewChat(nc)
}

func handleNewMessage(s *Server, ss *Session, e *Event) {
	var req NewMessageRequest
	if err := mapstructure.Decode(e.Data, &req); err != nil {
		fmt.Println(err)
		return
	}
	c, f := s.Chats[req.ChatID]
	if !f {
		fmt.Println("Message sent by user ", ss.User.ID, " to chat that doesn't exist ", req.ChatID)
		return
	}
	if _, f = c.Users[ss.User.ID]; !f {
		fmt.Println("Message sent by user ", ss.User.ID, " to chat that it isn't subscribed to ", req.ChatID)
	}
	resp := NewMessage(ss.User.ID, req.ChatID, req.Text)
	c.Broadcast <- NewEvent("new_message", resp)
}

func handleSetUser(s *Server, ss *Session, e *Event) {
	var req SetUserRequest
	err := mapstructure.Decode(e.Data, &req)
	if err != nil {
		fmt.Println(err)
		return
	}
	ss.User.Name = req.Name
	// BROADCAST LOGIC
}

func handleSubscribeChat(s *Server, ss *Session, e *Event) {
	var req ChatSubscribeRequest
	if err := mapstructure.Decode(e.Data, &req); err != nil {
		fmt.Println(err)
		return
	}
	if c, f := s.Chats[req.ChatID]; f {
		c.SubscribeUser(ss.User)
	}
}

func handleAuthentication(s *Server, ss *Session, e *Event) {
	var req AuthenticationRequest
	if err := mapstructure.Decode(e.Data, &req); err != nil {
		fmt.Println(err)
		return
	}
	if req.NewUser {
		u, err := NewUser()
		if err != nil {
			fmt.Println(err)
			return
		}
		t, err := GenerateToken(s, u)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = ss.Conn.WriteJSON(&Event{
			"authenticate",
			&AuthenticationResponse{u.ID, u.Name, t},
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		s.NewUser(u)
	} else {
		u, err := Authenticate(s, req.Token)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = ss.Conn.WriteJSON(&Event{
			"authenticate",
			&AuthenticationResponse{u.ID, u.Name, req.Token},
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

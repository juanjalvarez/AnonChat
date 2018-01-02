package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type EventHandler func(*Server, *Session, *Event)

type NewChatRequest struct {
	Name string `json:"name"`
}

type NewChatResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type SetUserRequest struct {
	UserID string `json:"userId"`
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
	s.Handle("new_chat", authTest(handleNewChat))
	s.Handle("new_message", authTest(handleNewMessage))
	s.Handle("set_user", authTest(handleSetUser))
	s.Handle("subscribe_chat", authTest(handleSubscribeChat))
	s.Handle("authenticate", handleAuthentication)
}

func authTest(eh EventHandler) EventHandler {
	return func(s *Server, ss *Session, e *Event) {
		if ss.User != nil {
			eh(s, ss, e)
		} else {
			ss.Send <- &Event{
				"failed_authentication",
				nil,
			}
		}
	}
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
	ss.User.RegisterChat(nc)
	ss.Send <- &Event{
		"chat_status",
		nc.GenerateStatus(s),
	}
}

func handleNewMessage(s *Server, ss *Session, e *Event) {
	fmt.Println(e.Type, "event triggered")
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
	evt := &Event{
		"set_user",
		&SetUserRequest{
			ss.User.ID,
			req.Name,
		},
	}
	fmt.Println("Changing name for", ss.User.UniqueIdentifier(), "to", req.Name)
	ss.User.Name = req.Name
	sentMap := make(map[string]bool)
	sentMap[ss.User.ID] = true
	ss.Send <- evt
	for _, c := range ss.User.Chats {
		for _, u := range c.Users {
			if _, sent := sentMap[u.ID]; !sent {
				if uss, fss := s.Sessions[u.ID]; fss {
					uss.Send <- evt
				}
				sentMap[u.ID] = true
			}
		}
	}
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
	var u *User
	var err error
	if req.NewUser {
		u, err = NewUser()
		if err != nil {
			fmt.Println(err)
			return
		}
		t, err := GenerateToken(s, u)
		if err != nil {
			fmt.Println(err)
			return
		}
		ss.User = u
		ss.Send <- &Event{
			"authenticate",
			&AuthenticationResponse{u.ID, u.Name, t},
		}
		s.NewUser(u)
	} else {
		u, err = Authenticate(s, req.Token)
		if err != nil {
			fmt.Println(err)
			return
		}
		ss.User = u
		ss.Send <- &Event{
			"authenticate",
			&AuthenticationResponse{u.ID, u.Name, req.Token},
		}
	}
	fmt.Println("User", u.UniqueIdentifier(), "authenticated")
	s.NewSession(ss)
	u.SendChats(s)
}

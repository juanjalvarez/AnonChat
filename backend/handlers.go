package main

import (
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

type MessageRequest struct {
	ChatID string `json:"chatId"`
	Text   string `json:"text"`
}

func registerHandlers(s *Server) {
	s.Handle("new_chat", authTest(handleNewChat))
	s.Handle("set_user", authTest(handleSetUser))
	s.Handle("join_chat", authTest(handleJoinChat))
	s.Handle("authenticate", handleAuthentication)
	s.Handle("message", authTest(handleMessage))
}

func authTest(eh EventHandler) EventHandler {
	return func(s *Server, ss *Session, e *Event) {
		if ss.User != nil {
			eh(s, ss, e)
		} else {
			ss.Send <- NewEvent("failed_authentication", nil)
		}
	}
}

func handleNewChat(s *Server, ss *Session, e *Event) {
	var req NewChatRequest
	if err := mapstructure.Decode(e.Data, &req); err != nil {
		s.log.err.Println(err)
		return
	}
	nc, err := NewChat(req.Name, ss.User)
	if err != nil {
		s.log.err.Println(err)
		return
	}
	s.NewChat(nc)
	ss.User.RegisterChat(nc)
	s.log.info.Println("Created chat", nc.UniqueIdentifier(), "by", ss.User.UniqueIdentifier())
	ss.Send <- NewEvent("chat_status", nc.GenerateStatus(s))
}

func handleMessage(s *Server, ss *Session, e *Event) {
	var req MessageRequest
	if err := mapstructure.Decode(e.Data, &req); err != nil {
		s.log.err.Println(err)
		return
	}
	c, f := s.Chats[req.ChatID]
	if !f {
		s.log.warn.Println("Message sent by user ", ss.User.ID, " to chat that doesn't exist ", req.ChatID)
		return
	}
	if _, f = c.Users[ss.User.ID]; !f {
		s.log.warn.Println("Message sent by user ", ss.User.ID, " to chat that it isn't subscribed to ", req.ChatID)
	}
	resp := NewMessage(ss.User.ID, req.ChatID, req.Text)
	c.Broadcast <- NewEvent("message", resp)
}

func handleSetUser(s *Server, ss *Session, e *Event) {
	var req SetUserRequest
	err := mapstructure.Decode(e.Data, &req)
	if err != nil {
		s.log.err.Println(err)
		return
	}
	evt := NewEvent("set_user", &SetUserRequest{
		ss.User.ID,
		req.Name,
	})
	s.log.info.Println("Changing name for", ss.User.UniqueIdentifier(), "to", req.Name)
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

func handleJoinChat(s *Server, ss *Session, e *Event) {
	var req ChatSubscribeRequest
	if err := mapstructure.Decode(e.Data, &req); err != nil {
		s.log.err.Println(err)
		return
	}
	if c, f := s.Chats[req.ChatID]; f {
		c.SubscribeUser(ss.User)
		ss.User.RegisterChat(c)
		s.log.info.Println("User", ss.User.UniqueIdentifier(), "joined chat", c.UniqueIdentifier())
		c.Broadcast <- NewEvent("chat_status", c.GenerateStatus(s))
	}
}

func handleAuthentication(s *Server, ss *Session, e *Event) {
	var req AuthenticationRequest
	if err := mapstructure.Decode(e.Data, &req); err != nil {
		s.log.err.Println(err)
		return
	}
	var u *User
	var t string
	var err error
	if req.NewUser {
		u, err = NewUser()
		if err != nil {
			s.log.err.Println(err)
			return
		}
		t, err = GenerateToken(s, u)
		if err != nil {
			s.log.err.Println(err)
			return
		}
		ss.User = u
		s.NewUser(u)
	} else {
		u, err = Authenticate(s, req.Token)
		if err != nil {
			s.log.err.Println(err)
			return
		}
		if u == nil {
			s.log.warn.Println("Failed to authenticate")
			return
		}
		t = req.Token
		ss.User = u
	}
	ss.Send <- NewEvent("authenticate", &AuthenticationResponse{u.ID, u.Name, t})
	s.log.info.Println("User", u.UniqueIdentifier(), "authenticated")
	s.NewSession(ss)
	u.SendChats(s)
}

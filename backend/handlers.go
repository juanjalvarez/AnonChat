package main

import (
	"errors"
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
	Text string `json:"text"`
}

type ChatSubscribeRequest struct {
	Chat string `json:"chat"`
}

func registerHandlers(s *Server) {
	s.Handle("new_chat", handleNewChat)
	s.Handle("new_message", handleNewMessage)
	s.Handle("set_user", handleSetUser)
	s.Handle("subscribe_chat", handleSubscribeChat)
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
	ss.Chat = nc
	// TODO
}

func handleNewMessage(s *Server, ss *Session, e *Event) {
	var req NewMessageRequest
	if err := mapstructure.Decode(e.Data, &req); err != nil {
		fmt.Println(err)
		return
	}
	if ss.Chat == nil {
		fmt.Println(errors.New("New message requested by a user that is not in any chat"))
		return
	}
	resp := NewMessage(ss.User.ID, req.Text)
	ss.Chat.Broadcast <- NewEvent("new_message", resp)
}

func handleSetUser(s *Server, ss *Session, e *Event) {
	var req SetUserRequest
	err := mapstructure.Decode(e.Data, &req)
	if err != nil {
		fmt.Println(err)
		return
	}
	ss.User.Name = req.Name
	if ss.Chat != nil {
		ss.Chat.Broadcast <- e
	}
}

func handleSubscribeChat(s *Server, ss *Session, e *Event) {
	var req ChatSubscribeRequest
	if err := mapstructure.Decode(e.Data, &req); err != nil {
		fmt.Println(err)
		return
	}
	if c, f := s.Chats[req.Chat]; f {
		ss.Chat = c
	}
}

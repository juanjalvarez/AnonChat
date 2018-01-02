package main

import "errors"

type User struct {
	ID    string
	Name  string
	Chats map[string]*Chat
}

type UserStatus struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Online bool   `json:"online"`
}

func NewUser() (*User, error) {
	newId, err := GenerateKey(8)
	if err != nil {
		return nil, err
	}
	return &User{
		string(newId),
		"anonymous",
		make(map[string]*Chat),
	}, nil
}

func (u *User) RegisterChat(c *Chat) {
	u.Chats[c.ID] = c
}

func (u *User) SendChats(s *Server) error {
	ss, f := s.Sessions[u.ID]
	if !f {
		return errors.New("The user has no active session")
	}
	for _, c := range u.Chats {
		ss.Send <- &Event{
			"chat_status",
			c.GenerateStatus(s),
		}
	}
	return nil
}

func (u *User) GenerateStatus(s *Server) *UserStatus {
	_, f := s.Sessions[u.ID]
	return &UserStatus{u.ID, u.Name, f}
}

func (u *User) UniqueIdentifier() string {
	return u.ID + "@" + u.Name
}

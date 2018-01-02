package main

import "sync"

type Chat struct {
	sync.Mutex
	ID        string
	Name      string
	Broadcast chan *Event
	Owner     *User
	Users     map[string]*User
}

type ChatStatus struct {
	ID    string                 `json:"id"`
	Name  string                 `json:"name"`
	Owner string                 `json:"owner"`
	Users map[string]*UserStatus `json:"users"`
}

func NewChat(name string, owner *User) (*Chat, error) {
	newID, err := GenerateKey(8)
	if err != nil {
		return nil, err
	}
	nc := &Chat{
		sync.Mutex{},
		string(newID),
		name,
		make(chan *Event),
		owner,
		make(map[string]*User),
	}
	if owner != nil {
		nc.SubscribeUser(owner)
	}
	return nc, nil
}

func (c *Chat) Write(s *Server) {
	for e := range c.Broadcast {
		for id := range c.Users {
			if ss, f := s.Sessions[id]; f {
				ss.Send <- e
			}
		}
	}
}

func (c *Chat) SubscribeUser(u *User) {
	c.Lock()
	c.Users[u.ID] = u
	c.Unlock()
}

func (c *Chat) GenerateStatus(s *Server) *ChatStatus {
	users := make(map[string]*UserStatus)
	for _, u := range c.Users {
		users[u.ID] = u.GenerateStatus(s)
	}
	return &ChatStatus{c.ID, c.Name, c.Owner.ID, users}
}

func (c *Chat) UniqueIdentifier() string {
	return c.ID + "@" + c.Name
}

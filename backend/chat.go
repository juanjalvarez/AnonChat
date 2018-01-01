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
		for id, _ := range c.Users {
			if ss, f := s.Sessions[id]; f {
				ss.Conn.WriteJSON(e)
			}
		}
	}
}

func (c *Chat) SubscribeUser(u *User) {
	c.Lock()
	c.Users[u.ID] = u
	c.Unlock()
	u.RegisterChat(c)
}

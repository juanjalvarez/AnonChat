package main

import "sync"

type Chat struct {
	sync.Mutex `json:"-"`
	ID         string `gorethink:"id",omitempty`
	Name       string `gorethink:"name"`
	Owner      string `gorethink:"ownerId"`
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
		owner.ID,
	}
	if owner != nil {
		nc.SubscribeUser(owner)
	}
	return nc, nil
}

func (c *Chat) SubscribeUser(u *User) {
	c.Lock()
	//c.Users[u.ID] = u
	c.Unlock()
}

func (c *Chat) GenerateStatus(s *Server) *ChatStatus {
	// users := make(map[string]*UserStatus)
	// for _, u := range c.Users {
	// 	users[u.ID] = u.GenerateStatus(s)
	// }
	// return &ChatStatus{c.ID, c.Name, c.Owner.ID, users}
	// TODO
	return nil
}

func (c *Chat) UniqueIdentifier() string {
	return c.ID + "@" + c.Name
}

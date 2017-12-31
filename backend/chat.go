package main

type Chat struct {
	ID        string
	Name      string
	Broadcast chan *Event
	Owner     *User
}

func NewChat(name string, owner *User) (*Chat, error) {
	newId, err := generateKey(8)
	if err != nil {
		return nil, err
	}
	return &Chat{
		string(newId),
		name,
		make(chan *Event),
		owner,
	}, nil
}

func (c *Chat) Write(s *Server) {
	for e := range c.Broadcast {
		for _, ss := range s.Sessions {
			if ss.Chat == c {
				ss.Conn.WriteJSON(e)
			}
		}
	}
}

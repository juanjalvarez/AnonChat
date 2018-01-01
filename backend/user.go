package main

type User struct {
	ID    string
	Name  string
	Chats map[string]*Chat
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

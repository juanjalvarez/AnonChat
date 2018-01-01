package main

type User struct {
	ID   string
	Name string
}

func NewUser() (*User, error) {
	newId, err := GenerateKey(8)
	if err != nil {
		return nil, err
	}
	return &User{
		string(newId),
		"anonymous",
	}, nil
}

package main

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	ID string `json:"id"`
}

func GenerateToken(s *Server, id string) (*Token, error) {
	_, f := s.Users[id]
	if !f {
		return nil, errors.New("User does not exist")
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	tokenString, err := t.SignedString(s.PrivateKey)
	if err != nil {
		return nil, err
	}
	return &Token{
		tokenString,
	}, nil
}

func Authenticate(s *Server, token *Token) (*User, error) {
	u, f := s.Users[token.ID]
	if !f {
		nu, err := NewUser()
		if err != nil {
			return nil, err
		}
		u = nu
	}
	return u, nil
}

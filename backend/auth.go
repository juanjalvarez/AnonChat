package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

type Claim struct {
	ID string `json:"id"`
}

func GenerateToken(s *Server, u *User) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": u.ID,
	})
	tokenString, err := t.SignedString(s.PrivateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Authenticate(s *Server, t string) (*User, error) {
	tok, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		return s.PrivateKey, nil
	})
	if err != nil {
		return nil, err
	}
	var c Claim
	err = mapstructure.Decode(tok.Claims, &c)
	if err != nil {
		return nil, err
	}
	u, f := s.Users[c.ID]
	if !f {
		return nil, nil
	}
	return u, nil
}

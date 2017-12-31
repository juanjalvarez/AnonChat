package main

import (
	"crypto/rand"
	"crypto/rsa"
	"io/ioutil"
)

func loadPrivateKey(path string) ([]byte, error) {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err != nil {
		key, err = generateKey(64)
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile(path, key, 0777)
		if err != nil {
			return nil, err
		}
	}
	return key, nil
}

func generateKey(len int) ([]byte, error) {
	key, err := rsa.GenerateKey(rand.Reader, 32*len)
	if err != nil {
		return nil, err
	}
	return []byte(key.D.Text(32)), nil
}

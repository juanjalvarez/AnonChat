package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	DB string `json:"db"`
}

func main() {
	l := NewLogger()
	var cfg Config
	d, err := ioutil.ReadFile("./config.json")
	if err != nil {
		l.err.Println(err)
		return
	}
	err = json.Unmarshal(d, &cfg)
	if err != nil {
		l.err.Println(err)
		return
	}
	srv, err := NewServer(&cfg, l)
	if err != nil {
		l.err.Println(err)
		return
	}
	mux := mux.NewRouter()
	mux.Handle("/", srv)
	http.Handle("/", mux)
	l.info.Println("Server running on port 4000")
	http.ListenAndServe(":4000", nil)
}

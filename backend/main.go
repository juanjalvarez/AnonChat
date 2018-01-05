package main

import (
	"net/http"

	r "github.com/dancannon/gorethink"
	"github.com/gorilla/mux"
)

func main() {
	l := NewLogger()
	_, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "rtsupport",
	})
	if err != nil {
		l.err.Println(err)
		return
	}
	srv, err := NewServer(l)
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

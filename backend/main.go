package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	srv, err := NewServer()
	if err != nil {
		fmt.Println(err)
		return
	}
	mux := mux.NewRouter()
	mux.Handle("/", srv)
	http.Handle("/", mux)
	http.ListenAndServe(":4000", nil)
}

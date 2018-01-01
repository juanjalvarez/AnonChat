package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	srv, err := NewServer()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	mux := mux.NewRouter()
	mux.Handle("/", srv)
	http.Handle("/", mux)
	fmt.Println("Server running on port 4000")
	http.ListenAndServe(":4000", nil)
}

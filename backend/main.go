package main

import (
	"fmt"
	"net/http"

	"github.com/googollee/go-socket.io"
	"github.com/rs/cors"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		fmt.Println("error:", err)
	}
	mux := http.NewServeMux()
	server.On("connection", func(so socketio.Socket) {
		fmt.Println("new connection")
		so.Join("chat")
		so.On("message", func(msg string) {
			so.Emit("message", msg)
			so.BroadcastTo("chat", msg)
			fmt.Println("received message:", msg)
		})
		so.On("disconnection", func() {
			fmt.Println("disconnected")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		fmt.Println("error:", err)
	})
	mux.Handle("/socketio", server)
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(cors.Default().Handler(mux))
	err = http.ListenAndServe(":5000", handler)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Listening on port 5000")
}

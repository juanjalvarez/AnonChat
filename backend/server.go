package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	sync.Mutex
	PrivateKey []byte
	Handlers   map[string]EventHandler
	Users      map[string]*User
	Sessions   map[string]*Session
	Chats      map[string]*Chat
}

func NewServer() (*Server, error) {
	abs, err := filepath.Abs("./private.key")
	if err != nil {
		return nil, err
	}
	key, err := loadPrivateKey(abs)
	if err != nil {
		return nil, err
	}
	s := &Server{
		sync.Mutex{},
		key,
		make(map[string]EventHandler),
		make(map[string]*User),
		make(map[string]*Session),
		make(map[string]*Chat),
	}
	registerHandlers(s)
	return s, nil
}

func (s *Server) Handle(event string, eh EventHandler) {
	s.Lock()
	s.Handlers[event] = eh
	s.Unlock()
}

func (s *Server) HandleEvent(ss *Session, e *Event) {
	if h, f := s.Handlers[e.Type]; f {
		h(s, ss, e)
	} else {
		fmt.Println("Failed to find event handler for", e.Type)
	}
}

func (s *Server) EndSession(ss *Session) {
}

func (s *Server) NewChat(chat *Chat) {
	s.Lock()
	s.Chats[chat.Name] = chat
	s.Unlock()
}

func (s *Server) NewUser(u *User) {
	s.Lock()
	s.Users[u.ID] = u
	s.Unlock()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	val := r.Header.Get("Authentication")
	fmt.Println(val)
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	ss := NewSession(socket)
	go ss.Read(s)
	go ss.Write(s)
}

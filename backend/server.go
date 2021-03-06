package main

import (
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
	log        *Logger
}

func NewServer(l *Logger) (*Server, error) {
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
		l,
	}
	registerHandlers(s)
	return s, nil
}

func (s *Server) Handle(event string, eh EventHandler) {
	s.Lock()
	s.Handlers[event] = eh
	s.Unlock()
	s.log.info.Println("Registered handler for '", event, "'")
}

func (s *Server) HandleEvent(ss *Session, e *Event) {
	if h, f := s.Handlers[e.Type]; f {
		h(s, ss, e)
	} else {
		s.log.warn.Println("Failed to find event handler for", e.Type)
	}
}

func (s *Server) NewSession(ss *Session) {
	s.Lock()
	s.Sessions[ss.User.ID] = ss
	s.Unlock()
	s.log.info.Println("New session for user", ss.User.UniqueIdentifier())
}

func (s *Server) EndSession(ss *Session) {
	if ss.User != nil {
		s.Lock()
		delete(s.Sessions, ss.User.ID)
		s.Unlock()
	}
	ss.Conn.Close()
	if ss.User != nil {
		s.log.info.Println("Terminating session for user", ss.User.UniqueIdentifier())
	} else {
		s.log.warn.Println("Terminating rogue sesssion")
	}
}

func (s *Server) NewChat(c *Chat) {
	s.Lock()
	s.Chats[c.ID] = c
	s.Unlock()
	go c.Write(s)
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
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.log.err.Println(err)
		return
	}
	ss := NewSession(socket)
	go ss.Read(s)
	go ss.Write(s)
}
